package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/cli"
	"github.com/mzahmi/ventilator/control/initialization"
	"github.com/mzahmi/ventilator/control/modeselect"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/params"
)

var UI = params.DefaultParams
var wg sync.WaitGroup

func main() {
	f, err := os.OpenFile("Events.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	//initialize the user input parameters
	params.InitParams(client)
	logger.Println("Parameters Initialized")

	wg.Add(5) //TODO: determine how to properly assign the number of goroutinues

	// Checks if GUI changed params and pushed to redis
	go func() {
		defer wg.Done()
		for {
			status, _ := client.Get("status").Result()
			readStatus <- status
			logger.Printf("Ventilation status changed to %s\n", status)
		}
	}()

	// Reads sensors and populate the graph
	go func() {
		defer wg.Done()
		for {
			Pin, Pout := sensors.ReadAllSensors()
			s <- sensors.SensorsReading{
				PressureInput:  Pin,
				PressureOutput: Pout}
			client.Set("pressure", Pin, 0).Err()
		}
	}()

	// Runs the ventelation method control
	go func() {
		defer wg.Done()
		for {
			for val := range readStatus {
				if val == "start" {
					go modeselect.ModeSelection(&UI, s, &wg, readStatus)
					client.Set("status", "ventilating", 0).Err()
					readStatus <- "ventilating"
					logger.Println("Ventilation status changed to ventilating")
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
	wg.Wait()
}

// prints out the checked error err
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
