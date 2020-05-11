package main

import (
	"fmt"
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
	//initialize the hardware
	initialization.HardwareInit()

	//establish connection with redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "dupi1.local:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	check(err)

	s := make(chan sensors.SensorsReading)
	readStatus := make(chan string)

	//initialize the user input parameters
	params.InitParams(client)

	wg.Add(5) //TODO: determine how to properly assign the number of gorotinues

	// Checks if GUI changed params and pushed to redis
	go func() {
		defer wg.Done()
		for {
			status, _ := client.Get("status").Result()
			readStatus <- status
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
					// write to redis status = ventilating
				} else if val == "stop" {
					// stop function to stop ventilation
				} else if val == "exit" {
					// exit program
					// exit <- true
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
