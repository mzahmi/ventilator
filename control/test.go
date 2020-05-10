package main

import (
	"fmt"
	"strconv"
	"time"
	"vent/control/initialization"
	"vent/control/modeselect"
	"vent/control/sensors"
	"vent/control/userinterface"

	"github.com/go-redis/redis"
)

var UI userinterface.UserInput

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
	start := make(chan bool)
	client.Set("start", "false", 0).Err()
	go func() {
		for {
			temp, _ := client.Get("IE").Result()
			temp1, _ := strconv.ParseFloat(temp, 32)
			UI.ER = float32(temp1)
			UI.IR = 1
			temp, _ = client.Get("BPM").Result()
			temp1, _ = strconv.ParseFloat(temp, 32)
			UI.Rate = float32(temp1)
			fmt.Println(UI.Rate)
			check, _ := client.Get("start").Result()
			if check == "true" {
				start <- true
				break
			}
		}
	}()
	go func() {
		for {
			Pin, Pout := sensors.ReadAllSensors()
			s <- sensors.SensorsReading{
				PressureInput:  Pin,
				PressureOutput: Pout}
			client.Set("pressure", Pin, 0).Err()
		}
	}()
	for {
		select {
		case <-start:
			go modeselect.PressureControl(UI, s)
		case <-time.After(240 * time.Second):
			fmt.Println("about to exit program")
			return
		}
	}
}
