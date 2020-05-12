package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/cli"
	"github.com/mzahmi/ventilator/control/initialization"
	"github.com/mzahmi/ventilator/control/modeselect"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
	"github.com/mzahmi/ventilator/params"
)

var UI = params.DefaultParams
var wg sync.WaitGroup

func main() {
	f, err := os.OpenFile("Events.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	// declare new logger
	logger := log.New(f, "Event", log.LstdFlags)

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
	check(err)

	//delcare channels to communicate between goroutines
	s := make(chan sensors.SensorsReading)
	readStatus := make(chan string)
	client.Set("status", "NA", 0).Err() // empty the previous record of status

	//initialize the user input parameters
	params.InitParams(client)
	logger.Println("Parameters Initialized")

	//TODO: determine how to properly assign the number of goroutinues
	wg.Add(5)

	// Checks if GUI changed params and pushed to redis
	go func() {
		defer wg.Done()
		for {
			status, _ := client.Get("status").Result()
			readStatus <- status
		}
	}()

	// Reads sensors and populate the graph
	//limit the reading frequency to a predefined value
	rate := float64(100)                                                // Hz rate
	timePerLoopIteration := time.Duration(1000/rate) * time.Millisecond //(1 / rate) ms
	fmt.Println(timePerLoopIteration)

	go func() {
		defer wg.Done()
		for {
			t1 := time.Now()

			Pin, Pout := sensors.ReadAllSensors()
			s <- sensors.SensorsReading{
				PressureInput:  Pin,
				PressureOutput: Pout}
			client.Set("pressure", Pin*1020, 0).Err()

			loopTime := time.Since(t1)
			//fmt.Println("Loop time:", loopTime)
			if loopTime < timePerLoopIteration {
				diff := (timePerLoopIteration - loopTime)
				//fmt.Println("Sleeping for:", diff)
				time.Sleep(diff)
			}
			t3 := time.Now()
			fmt.Println("Tdiff=", t3.Sub(t1))
		}
	}()

	// Runs the ventelation method control
	go func() {
		defer wg.Done()
		for {
			for val := range readStatus {
				if val == "start" {
					logger.Printf("Ventilation status changed to %s\n", val)
					UI = params.ReadParams(client)
					go modeselect.ModeSelection(&UI, s, &wg, readStatus, logger)
					client.Set("status", "ventilating", 0).Err()
					readStatus <- "ventilating"
					logger.Printf("Ventilation status changed to %v", <-readStatus)
					// write to redis status = ventilating
				} else if val == "stop" {
					// stop function to stop ventilation
					//fmt.Println("Stopping system")
					logger.Println("Stopping system")
				} else if val == "exit" {
					// exit program
					// exit <- true
					logger.Println("Exiting system")
				}
			}
		}
	}()

	// Provides CLI interface
	go cli.Run(&wg, s, client, readStatus)
	SetupCloseHandler()
	wg.Wait()
}

// prints out the checked error err
func check(err error) {
	if err != nil {
		fmt.Println(err)
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
