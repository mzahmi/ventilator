// Package modeselect ...
package modeselect

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
	"github.com/mzahmi/ventilator/logger"
	"github.com/mzahmi/ventilator/params"
)

// VolumeAC one of the main 5 modes of the github.com/mzahmi/ventilatorilator
func VolumeAC(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logStruct *logger.Logging) {
	switch UI.BreathType {
	case "Volume Control":
		VolumeControl(UI, s, client, mux, logStruct)
	case "Volume Assist":
		VolumeAssist(UI, s, client, mux, logStruct)
	default:
		fmt.Println("Enter valid breath type")
	}
}

// VolumeControl a submode of Volume AC
// 	Triggering:	Time
// 	Cycling: 	Time
// 	Control: 	Volume
func VolumeControl(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logStruct *logger.Logging) {

	// UI.PeakFlow = UI.TidalVolume / UI.Ti
	// //initiate a PID controller based on the PeakFlow
	// FlowPID := NewPIDController(0.5, 0.5, 0) // takes in P, I, and D values
	// FlowPID.SetOutputLimits(0, 0.5)
	// // FlowPID.Set(float64(UI.TidalVolume / 1000))
	// FlowPID.Set(float64(UI.PeakFlow))
	// fmt.Printf("recieved tidal volume: %v\n", UI.TidalVolume)
	desireFlowRate:= (UI.TidalVolume*60)/(UI.Ti*1000)
	desireVoltage:= desireFlowRate/100
	// fmt.Printf("desired voltage : %v\n", desireVoltage)
	

	//control loop
	for {
		//FlowPID.setpoint = float64(UI.PeakFlow) // Sets the PID setpoint
		//Open main valve MIns controlled by flow sensor PIns
		var volume float32
		valves.MV.Open()
		valves.MExp.Close()
		valves.InProp.IncrementValve(float64(desireVoltage))
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
			tic:= time.Now()
			mux.Lock()
			flowRate := s.FlowInput
			mux.Unlock()
			runtime.Gosched()
			//valves.InProp.IncrementValve(FlowPID.Update(float64(incrementor)))
			time.Sleep(1 * time.Millisecond)
			toc:= time.Since(tic)
			volume += float32(toc.Minutes()) * flowRate
			client.Set("volume", volume, 0).Err()
		}
		//Close main valve MIns
		client.Set("volume", 0, 0).Err()
		valves.InProp.Close() // closes the valve
		valves.InProp.Close() // closes the valve
		// time.Sleep(1*time.Millisecond)
		//Open main valve MExp controlled by flow sensor PExp
		// for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
		// 	if sensors.PExp.ReadPressure() <= UI.PEEP {
		// 		break
		// 	}
		// 	valves.ExProp.IncrementValve(100)
		// }
		valves.MExp.Open()
		time.Sleep(time.Duration(UI.Te*1000) * time.Millisecond)
		//Close main valve MExp
		valves.MExp.Close() // closes the valve
		// if it's stop or exit then close valves and break loop
		trig, err := client.Get("status").Result()
		check(err, logStruct)
		if (trig == "stop") || (trig == "exit") {
			break
		} else {
			continue
		}
	}
	valves.CloseAllValves(&valves.MV, &valves.MExp, &valves.InProp)
	// logger.Println("All valves closed")
	logStruct.Event("All valves closed")
}

// VolumeAssist a submode of Volume AC. The Triggering type is chosen by the operator.
// 	Triggering:	Pressure/Flow
// 	Cycling: 	Time
// 	Control: 	Volume
func VolumeAssist(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logStruct *logger.Logging) {

	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//Check trigger type
	switch UI.PatientTriggerType {
	case "Pressure Trigger":
		logStruct.Event(fmt.Sprintf("Pressure Trigger is set at %v cmH2O\n", UI.PressureTrigSense))
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
			check(err, logStruct)
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
		logStruct.Event(fmt.Sprintf("Flow Trigger is set at %v cmH2O\n", UI.FlowTrigSense))
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
			check(err, logStruct)
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
