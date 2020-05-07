package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/gomodule/redigo/redis"
)

type UserInput struct {
	Mode                string
	BreathType          string
	PatientTriggerType  string
	TidalVolume         float32 // ml
	Rate                float32 // BPM
	Ti                  float32 // inhalation time
	TiMax               float32 // for PSV mode backup time control
	Te                  float32 // exhalation time
	IR                  float32 // inhalation ratio part IR:IE
	ER                  float32 // exhalation ratio part IR:IE
	PeakFlow            float32
	PEEP                float32 // 5-20 mmH2O
	FiO2                float32 // 21% - 100%
	PressureTrigSense   float32 // -0.5 to 02 mmH2O
	FlowTrigSense       float32 // 0.5 to 5 Lpm
	FlowCyclePercent    float32 // for flow cycling ranges from 0 to 100%
	PressureSupport     float32 // needs to be defined
	InspiratoryPressure float32 // Also known as P_control
	UpperLimitVT        float32 // upper limit of tidal volume
	LowerLimitVt        float32 // lower limit of tidal volume
	RiseTime            float32 // needs to be defined
	UpperLimitPIP       float32 // upper limit of airway peak prssure
	LowerLimitPIP       float32 // lower limit of airway peak pressure
	MinuteVolume        float32 // minute volume in Lpm is the amount of gas expired per minute.
	UpperLimitMV        float32 // upper limit of minute volume
	LowerLimitMV        float32 // lower limit of minute volume
	UpperLimitRR        float32 // upper limit of monitored BPM
	LowerLimitRR        float32 // lower limit of monitored BPM
}

var UI = UserInput{
	Mode:                "NA",
	BreathType:          "NA",
	PatientTriggerType:  "NA",
	TidalVolume:         0,
	Rate:                0,
	Ti:                  0,
	TiMax:               0,
	Te:                  0,
	IR:                  0,
	ER:                  0,
	PeakFlow:            0,
	PEEP:                0,
	FiO2:                0,
	PressureTrigSense:   0,
	FlowTrigSense:       0,
	FlowCyclePercent:    0,
	PressureSupport:     0,
	InspiratoryPressure: 0,
	UpperLimitVT:        0,
	LowerLimitVt:        0,
	RiseTime:            0,
	UpperLimitPIP:       0,
	LowerLimitPIP:       0,
	MinuteVolume:        0,
	UpperLimitMV:        0,
	LowerLimitMV:        0,
	UpperLimitRR:        0,
	LowerLimitRR:        0,
}

func info() {
	fmt.Println("CLI v1.0 Info:\n")
	fmt.Println("	info: i")
	fmt.Println("	quit: q")

	fmt.Println("	read: r sensor_name")
	fmt.Println("	read: rp parameter_name")
	fmt.Println("	write: w o_name value")
	fmt.Println("	write: wp parameter_name value")
	fmt.Println("	list sensors: lss")
	fmt.Println("	list actuators: lsa")
	fmt.Println("	list parameters: lsp")
}

func parameterPublisher(c chan UserInput) {
	for i := 0; ; i++ {
		UI.UpperLimitMV = float32(i)
		c <- UI
		time.Sleep(1 * time.Second)
	}
}

func cli(c chan UserInput) {

	conn, err := redis.Dial("tcp", "dupi1.local:6379")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Do("SET", "io:pressure", 100)
	if err != nil {
		fmt.Print("Error")
	}

	reply, err := redis.String(conn.Do("GET", "io:pressure"))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

	parameters := structs.Names(&UserInput{})

	fmt.Println("Type (i) for more info \n")
	reader := bufio.NewReader(os.Stdin)
	for {
		//user_input := <-c
		fmt.Print("> ")

		// get user input
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(input)

		// show information
		if words[0] == "i" {
			info()
			continue
		}

		// check for write
		if words[0] == "w" {
			if len(words) < 3 {
				fmt.Println("Incorrect number of parameters")
				continue
			}

			writeto := words[1]
			writevalue := words[2]
			fmt.Println("Changing actuator " + writeto + " to " + writevalue)
			continue
		}

		if words[0] == "wp" {
			if len(words) < 3 {
				fmt.Println("Incorrect number of parameters")
				continue
			}

			writeto := words[1]
			writevalue := words[2]
			fmt.Println("Changing parameter " + writeto + " to " + writevalue)
			continue
		}

		if words[0] == "lsp" {
			fmt.Println("Displaying the list of parameters")
			fmt.Println(strings.Join(parameters, "\n"))
			continue
		}

		if words[0] == "lss" {
			fmt.Println("Displaying the list of sensors")
			fmt.Println(strings.Join(parameters, "\n"))
			continue
		}

		if words[0] == "lsa" {
			fmt.Println("Displaying the list of actuators")
			fmt.Println(strings.Join(parameters, "\n"))
			continue
		}

		if words[0] == "r" {
			/*
				switch words[1]; {
				case "":

				}
			*/
			continue
		}

		if words[0] == "rp" {
			/*
				switch words[1]; {
				case "":

				}
			*/
			continue
		}

		if words[0] == "q" {
			break
		}
		fmt.Println("Unknown input")
	}
	defer conn.Close()
}

func main() {
	ch := make(chan UserInput)
	go parameterPublisher(ch)
	go cli(ch)
	for {

	}
}
