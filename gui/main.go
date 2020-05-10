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
	initialization.HardwareInit()
	//establish connection with redis
	client := redis.NewClient(&redis.Options{
		Addr:     "dupi1.local:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}

	s := make(chan sensors.SensorsReading)
	//exit := make(chan bool)
	//start := make(chan bool)
	readStatus := make(chan string)

	client.Set("start", "false", 0).Err() // clears previous entry of the start in redis

	params.InitParams(client)

	// Checks if GUI changed params and pushed to redis
	go func() {
		defer wg.Done()
		for {
			status, _ := client.Get("status").Result()
			readStatus <- status
		}
	}()

	// Reads sensors and share
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
	wg.Add(5)
	go cli.Run(&wg, s, client, readStatus)
	wg.Wait()
}
