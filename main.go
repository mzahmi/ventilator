package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/cli"
	"github.com/mzahmi/ventilator/control/initialization"
	"github.com/mzahmi/ventilator/control/modeselect"
	"github.com/mzahmi/ventilator/control/rpigpio"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
	"github.com/mzahmi/ventilator/params"
	// "github.com/mzahmi/ventilator/control/alarms"
)

var UI = params.DefaultParams
var wg sync.WaitGroup
var mux sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//Creates log file called Events.log
	f, err := os.OpenFile("file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// declare new loggers with different prefixes
	logger := log.New(f, "Event ", log.LstdFlags) // event logger
	logErr := log.New(f, "Error ", log.LstdFlags) // errors logger
	logArm := log.New(f, "Alarm ", log.LstdFlags) // alarms logger

	//initialize the hardware
	initialization.HardwareInit()
	logger.Println("Hardware Initialized")

	//establish connection with redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "dupi1.local:6379",
		Password: "",
		DB:       0,
	})
	_, err = client.Ping().Result()
	check(err, logErr)
	logger.Println("Client Initialized")

	// set the critical records in redis to zero or NA
	client.Set("status", "NA", 0).Err()
	client.Set("pressure", 0, 0).Err()
	client.Set("volume", 0, 0).Err()
	client.Set("flow", 0, 0).Err()

	//initialize the user input parameters
	params.InitParams(client)
	logger.Println("Parameters Initialized")

	//initialize a SensorsReading struct to store all of the sensor readings
	s := sensors.SensorsReading{
		PressureInput:  0,
		PressureOutput: 0,
	}

	// Reads sensors and populate the graph
	// limit the reading frequency to a predefined value
	rate := float64(500)                                                   // Hz rate
	timePerLoopIteration := time.Duration(1000000/rate) * time.Microsecond //(1 / rate) us

	go func() {
		for {
			t1 := time.Now()
			// reads all of the sensors on the system
			Pin, Pout := sensors.ReadAllSensors()
			//locks the populating of the sensors struct
			mux.Lock()
			s = sensors.SensorsReading{
				PressureInput:  Pin,
				PressureOutput: Pout}
			mux.Unlock()
			runtime.Gosched()
			//sends the pressure reading from Pin to GUI
			client.Set("pressure", (Pin)*1020, 0).Err()
			client.Set("volume", (Pin)*1020, 0).Err()
			client.Set("flow", (Pin)*1020, 0).Err()
			//fmt.Println(Pin*1020)
			//calculates the delay based on a specified rate
			loopTime := time.Since(t1)
			if loopTime < timePerLoopIteration {
				diff := (timePerLoopIteration - loopTime)
				//fmt.Println("Sleeping for:", diff)
				time.Sleep(diff)
			}
			//t3 := time.Now()
			//fmt.Println("Tdiff=", t3.Sub(t1))
		}
	}()

	//Airway pressure alarm check
	go func() {
		for {
			mux.Lock()
			airpress := s.PressureInput
			mux.Unlock()
			runtime.Gosched()
			if (airpress * 1020) >= 80 {
				//msg := "Airway Pressure high"
				client.Set("alarm_status", "critical", 0).Err()
				client.Set("alarm_title", "Airway Pressure high", 0).Err()
				client.Set("alarm_text", "Airway Pressure exceeded limits check for obstruction", 0).Err()
				logArm.Println("Airway Pressure high")
				tm := 200 * time.Millisecond
				ts := 3000 * time.Millisecond

				err := rpigpio.BeepOn()
				check(err, logErr)
				time.Sleep(tm)
				err = rpigpio.BeepOff()
				check(err, logErr)
				time.Sleep(tm)
				err = rpigpio.BeepOn()
				check(err, logErr)
				time.Sleep(tm)
				err = rpigpio.BeepOff()
				check(err, logErr)
				time.Sleep(tm)
				err = rpigpio.BeepOn()
				check(err, logErr)
				time.Sleep(tm)
				err = rpigpio.BeepOff()
				check(err, logErr)
				time.Sleep(tm)
				time.Sleep(ts)

			} else {
				client.Set("alarm_status", "none", 0).Err()
			}
			time.Sleep(time.Millisecond * 100)
		}

	}()

	// Provides CLI interface
	go cli.Run(&s, client, &mux)

	//checks for sys interupt
	SetupCloseHandler()

	for {
		status, err := client.Get("status").Result()
		check(err, logErr)
		if status == "start" {
			logger.Printf("Ventilation status changed to %s\n", status)
			UI = params.ReadParams(client)
			go modeselect.ModeSelection(&UI, &s, client, &mux, logger, logErr)
			client.Set("status", "ventilating", 0).Err()
			logger.Printf("Ventilation status changed to %s\n", status)
		} else if status == "stop" {
			logger.Println("Stopping system")
		} else if status == "exit" {
			logger.Println("Exiting system")
		}
	}
}

// prints out the checked error err
func check(err error, logErr *log.Logger) {
	if err != nil {
		logErr.Println(err)
	}
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		valves.CloseAllValves(&valves.InProp, &valves.MExp, &valves.MV)
		os.Exit(0)
	}()
}
