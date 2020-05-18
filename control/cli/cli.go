package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/structs"
	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/params"
)

func info() {
	fmt.Printf("CLI v1.0 Info:\n")
	fmt.Println("	info: i")
	fmt.Println("	quit: q")

	fmt.Println("	read: r sensor_name")
	fmt.Println("	read: rp parameter_name")
	fmt.Println("	write: w o_name value")
	fmt.Println("	write: wp parameter_name value")
	fmt.Println("	list sensors: lss")
	fmt.Println("	list actuators: lsa")
	fmt.Println("	list parameters: lsp")
	fmt.Println("	start ventilation: vstart")
	fmt.Println("	stop ventilation: vstop")
}

func Run(s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex) {

	parameters := structs.Names(&params.UserInput{})

	fmt.Printf("Type (i) for more info \n")
	reader := bufio.NewReader(os.Stdin)
	for {
		//user_input := <-c
		status, err := client.Get("status").Result()
		check(err)
		if status == "exit" {
			break
		}
		fmt.Print("> ")

		// get user input
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(input)
		if len(words) == 0 {
			continue
		}

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
			//TODO: write them to redis
			parameters := params.ReadParams(client)
			switch writeto {
			case "Mode":
				//TODO: check that values are correct
				parameters.Mode = writevalue
			case "BreathType":
				//TODO: check that values are correct
				parameters.BreathType = writevalue
			case "PatientTriggerType":
				//TODO: check that values are correct
				parameters.PatientTriggerType = writevalue
			case "TidalVolume":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.TidalVolume = float32(temp)
			case "Rate":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.Rate = float32(temp)
			case "Ti":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.Ti = float32(temp)
			case "TiMax":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.TiMax = float32(temp)
			case "Te":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.Te = float32(temp)
			case "IR":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.IR = float32(temp)
			case "ER":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.ER = float32(temp)
			case "PeakFlow":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.PeakFlow = float32(temp)
			case "PEEP":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.PEEP = float32(temp)
			case "FiO2":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.FiO2 = float32(temp)
			case "PressureTrigSense":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.PressureTrigSense = float32(temp)
			case "FlowTrigSense":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.FlowTrigSense = float32(temp)
			case "FlowCyclePercent":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.FlowCyclePercent = float32(temp)
			case "PressureSupport":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.PressureSupport = float32(temp)
			case "InspiratoryPressure":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.InspiratoryPressure = float32(temp)
			case "UpperLimitVT":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.UpperLimitVT = float32(temp)
			case "LowerLimitVT":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.LowerLimitVT = float32(temp)
			case "RiseTime":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.RiseTime = float32(temp)
			case "UpperLimitPIP":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.UpperLimitPIP = float32(temp)
			case "LowerLimitPIP":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.LowerLimitPIP = float32(temp)
			case "MinuteVolume":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.MinuteVolume = float32(temp)
			case "UpperLimitMV":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.UpperLimitMV = float32(temp)
			case "LowerLimitMV":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.LowerLimitMV = float32(temp)
			case "UpperLimitRR":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.UpperLimitRR = float32(temp)
			case "LowerLimitRR":
				temp, _ := strconv.ParseFloat(writevalue, 32)
				parameters.LowerLimitRR = float32(temp)

			}
			params.WriteParams(client, parameters)
			continue
		}

		if words[0] == "lsp" {
			fmt.Println("Displaying the list of parameters")
			fmt.Println(strings.Join(parameters, "\t\n"))
			continue
		}

		if words[0] == "lss" {
			fmt.Println("Displaying the list of sensors")
			fmt.Println("	PIns")
			fmt.Println("	PExp")
			fmt.Println("	FInsp")
			fmt.Println("	FExp")
			continue
		}

		if words[0] == "lsa" {
			fmt.Println("Displaying the list of actuators")
			fmt.Println(strings.Join(parameters, "\n"))
			continue
		}
		if words[0] == "vstart" {
			fmt.Println("Starting ventilation")
			client.Set("status", "start", 0).Err()
			continue
		}
		if words[0] == "vstop" {
			fmt.Println("Stop ventilation")
			client.Set("status", "stop", 0).Err()
			continue
		}

		if words[0] == "r" {
			if len(words) < 2 {
				fmt.Println("Incorrect number or arguments, press i for help")
				continue
			}
			switch words[1] {
			case "PIns":
				mux.Lock()
				val := s.PressureInput
				mux.Unlock()
				runtime.Gosched()
				fmt.Println("PIns:", val)
			case "PExp":
				mux.Lock()
				val := s.PressureOutput
				mux.Unlock()
				runtime.Gosched()
				fmt.Println("PExp:", val)
			case "FIns":
				//val := sensors.FIns.ReadFlow()
				val := 0
				fmt.Println("FIns:", val)
			case "FExp":
				//val := sensors.FExp.ReadFlow()
				val := 0
				fmt.Println("FExp:", val)
			default:
				fmt.Println("Unknow sensors")
			}
			continue
		}

		if words[0] == "rp" {
			if len(words) < 2 {
				fmt.Println("Incorrect number or arguments, press i for help")
				continue
			}

			parameters := params.ReadParams(client)
			switch words[1] {
			case "Mode":
				//TODO: check that values are correct
				fmt.Println("Mode:", parameters.Mode)
			case "BreathType":
				//TODO: check that values are correct
				fmt.Println("BreathType:", parameters.BreathType)
			case "PatientTriggerType":
				//TODO: check that values are correct
				fmt.Println("PatientTriggerType:", parameters.PatientTriggerType)
			case "TidalVolume":
				fmt.Println("TidalVolume:", parameters.TidalVolume)
			case "Rate":
				fmt.Println("Rate:", parameters.Rate)
			case "Ti":
				fmt.Println("Ti:", parameters.Ti)
			case "TiMax":
				fmt.Println("TiMax:", parameters.TiMax)
			case "Te":
				fmt.Println("Te:", parameters.Te)
			case "IR":
				fmt.Println("IR:", parameters.IR)
			case "ER":
				fmt.Println("ER:", parameters.ER)
			case "PeakFlow":
				fmt.Println("PeakFlow:", parameters.PeakFlow)
			case "PEEP":
				fmt.Println("PEEP:", parameters.PEEP)
			case "FiO2":
				fmt.Println("FiO2:", parameters.FiO2)
			case "PressureTrigSense":
				fmt.Println("PressureTrigSense:", parameters.PressureTrigSense)
			case "FlowTrigSense":
				fmt.Println("FlowTrigSense:", parameters.FlowTrigSense)
			case "FlowCyclePercent":
				fmt.Println("FlowCyclePercent:", parameters.FlowCyclePercent)
			case "PressureSupport":
				fmt.Println("PressureSupport:", parameters.PressureSupport)
			case "InspiratoryPressure":
				fmt.Println("InspiratoryPressure:", parameters.InspiratoryPressure)
			case "UpperLimitVT":
				fmt.Println("UpperLimitVT:", parameters.UpperLimitVT)
			case "LowerLimitVT":
				fmt.Println("LowerLimitVT:", parameters.LowerLimitVT)
			case "RiseTime":
				fmt.Println("RiseTime:", parameters.RiseTime)
			case "UpperLimitPIP":
				fmt.Println("UpperLimitPIP:", parameters.UpperLimitPIP)
			case "LowerLimitPIP":
				fmt.Println("LowerLimitPIP:", parameters.LowerLimitPIP)
			case "MinuteVolume":
				fmt.Println("MinuteVolume:", parameters.MinuteVolume)
			case "UpperLimitMV":
				fmt.Println("UpperLimitMV:", parameters.UpperLimitMV)
			case "LowerLimitMV":
				fmt.Println("LowerLimitMV:", parameters.LowerLimitMV)
			case "UpperLimitRR":
				fmt.Println("UpperLimitRR:", parameters.UpperLimitRR)
			case "LowerLimitRR":
				fmt.Println("LowerLimitRR:", parameters.LowerLimitRR)
			case "status":
				val, _ := client.Get("status").Result()
				fmt.Println("status:", val)
			default:
				fmt.Println("Invalid parameter")

			}
			continue
		}

		if words[0] == "q" {
			os.Exit(0)
			//break
		}
		fmt.Println("Unknown input")
	}
}

// prints out the checked error err
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
