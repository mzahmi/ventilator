// Package modeselect ...
package modeselect

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
	"github.com/mzahmi/ventilator/params"
)

// VolumeAC one of the main 5 modes of the github.com/mzahmi/ventilatorilator
func VolumeAC(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logger *log.Logger) {
	switch UI.BreathType {
	case "Volume Control":
		VolumeControl(UI, s, client, mux, logger)
	case "Volume Assist":
		VolumeAssist(UI, s, client, mux, logger)
	default:
		fmt.Println("Enter valid breath type")
	}
}

// VolumeControl a submode of Volume AC
// 	Triggering:	Time
// 	Cycling: 	Time
// 	Control: 	Volume
func VolumeControl(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logger *log.Logger) {

	//initiate a PID controller based on the PeakFlow
	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//control loop
	for {
		FlowPID.setpoint = float64(UI.PeakFlow) // Sets the PID setpoint
		//Open main valve MIns controlled by flow sensor PIns
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
			valves.InProp.IncrementValve(FlowPID.Update(float64(sensors.FIns.ReadFlow())))
		}
		//Close main valve MIns
		valves.InProp.IncrementValve(0) // closes the valve
		//Open main valve MExp controlled by flow sensor PExp
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
			if sensors.PExp.ReadPressure() <= UI.PEEP {
				break
			}
			valves.ExProp.IncrementValve(100)
		}
		//Close main valve MExp
		valves.ExProp.IncrementValve(0) // closes the valve
		// if it's stop or exit then close valves and break loop
		trig, err := client.Get("status").Result()
		check(err)
		if (trig == "stop") || (trig == "exit") {
			// valves.CloseAllValves(&valves.MIns, &valves.MExp)
			// logger.Println("All valves closed")
			break
		} else {
			continue
		}
	}
}

// VolumeAssist a submode of Volume AC. The Triggering type is chosen by the operator.
// 	Triggering:	Pressure/Flow
// 	Cycling: 	Time
// 	Control: 	Volume
func VolumeAssist(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logger *log.Logger) {

	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//Check trigger type
	switch UI.PatientTriggerType {
	case "Pressure Trigger":
		//Calculate trigger threshhold with PEEP and sensitivity
		PTrigger := UI.PEEP + UI.PressureTrigSense
		//Begin loop
		for {
			//check if trigger is true
			if sensors.PIns.ReadPressure() <= PTrigger {
				FlowPID.setpoint = float64(UI.PeakFlow) // Sets PID setpoint to PeakFlow
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
					valves.InProp.IncrementValve(FlowPID.Update(float64(sensors.FIns.ReadFlow())))
				}
				//Close main valve MIns
				valves.InProp.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled by flow sensor FExp
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te) * time.Second); {
					valves.ExProp.IncrementValve(FlowPID.Update(float64(sensors.FExp.ReadFlow())))
				}
				//Close main valve MExp
				valves.ExProp.IncrementValve(0) // closes the valve
			}
			// if it's stop or exit then close valves and break loop
			trig, err := client.Get("status").Result()
			check(err)
			if (trig == "stop") || (trig == "exit") {
				// valves.CloseAllValves(&valves.MIns, &valves.MExp)
				// logger.Println("All valves closed")
				break
			} else {
				continue
			}
		}
	case "Flow Trigger":
		//Calculate trigger threshhold with PEEP and sensitivity
		FTrigger := UI.FlowTrigSense
		//Begin loop
		for {
			//check if trigger is true
			if sensors.PIns.ReadPressure() >= FTrigger {
				FlowPID.setpoint = float64(UI.PeakFlow) // Sets PID setpoint to PeakFlow
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
					valves.InProp.IncrementValve(FlowPID.Update(float64(sensors.FIns.ReadFlow())))
				}
				//Close main valve MIns
				valves.InProp.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled by flow sensor FExp
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te) * time.Second); {
					valves.ExProp.IncrementValve(FlowPID.Update(float64(sensors.FExp.ReadFlow())))
				}
				//Close main valve MExp
				valves.ExProp.IncrementValve(0) // closes the valve
			}
			// if it's stop or exit then close valves and break loop
			trig, err := client.Get("status").Result()
			check(err)
			if (trig == "stop") || (trig == "exit") {
				// valves.CloseAllValves(&valves.MIns, &valves.MExp)
				// logger.Println("All valves closed")
				break
			} else {
				continue
			}
		}

	}
}
