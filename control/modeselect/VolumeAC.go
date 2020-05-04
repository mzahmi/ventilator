// Package modeselect Deals with the Volume AC
package modeselect

import (
	"fmt"
	"time"

	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
)

// VolumeAC one of the main 5 modes of the ventilator
func VolumeAC(UI *UserInput) {
	switch UI.BreathType {
	case "Control":
		VolumeControl(UI)
	case "Assist":
		VolumeAssist(UI)
	default:
		fmt.Println("Enter valid breath type")
	}
}

// VolumeControl a submode of Volume AC
// 	Triggering:	Time
// 	Cycling: 	Time
// 	Control: 	Volume
func VolumeControl(UI *UserInput) {

	//initiate a PID controller based on the PeakFlow
	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//control loop
	for !Exit {
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
	}

}

// VolumeAssist a submode of Volume AC. The Triggering type is chosen by the operator.
// 	Triggering:	Pressure/Flow
// 	Cycling: 	Time
// 	Control: 	Volume
func VolumeAssist(UI *UserInput) {

	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//Check trigger type
	switch UI.PatientTriggerType {
	case "Pressure":
		//Calculate trigger threshhold with PEEP and sensitivity
		PTrigger := UI.PEEP + UI.PressureTrigSense
		//Begin loop
		for !Exit {
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
		}
	case "Flow":
		//Calculate trigger threshhold with PEEP and sensitivity
		FTrigger := UI.FlowTrigSense
		//Begin loop
		for !Exit {
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
		}

	}
}
