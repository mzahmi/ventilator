package modeselect

import (
	"time"

	"vent/control/sensors"
	"vent/control/valves"
)

// PSV one of the main 5 modes of the ventilator. The Triggering type is chosen by the operator.
// 	Triggering:	Pressure/Flow
// 	Cycling: 	Flow
// 	Control: 	Pressure
func PSV(UI *UserInput) {

	PressurePID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//Check trigger type
	switch UI.PatientTriggerType {
	case "Pressure":
		//Calculate trigger threshhold with PEEP and sensitivity
		PTrigger := UI.PEEP + UI.PressureTrigSense

		counter := 1
		//Begin loop
		for !Exit {
			FlowMonitor := []float32{sensors.FIns.ReadFlow()}
			//check if trigger is true
			if sensors.PIns.ReadPressure() <= PTrigger {
				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.TiMax*1000) * time.Millisecond); {
					valves.InProp.IncrementValve(PressurePID.Update(float64(sensors.PIns.ReadPressure())))
					FlowMonitor = append(FlowMonitor, sensors.FIns.ReadFlow())
					if FlowMonitor[0] < FlowMonitor[counter] {
						FlowMonitor[0] = FlowMonitor[counter] // saves maximum value @ index 0
					} else if (FlowMonitor[0]*UI.FlowCyclePercent)/100 >= sensors.FIns.ReadFlow() {
						counter = 1
						FlowMonitor = nil
						break
					}
					counter++
				}
				//Close main valve MIns
				valves.InProp.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled by flow sensor FExp
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
	case "Flow":
		//Calculate trigger threshhold with flow trig sensitivity
		FTrigger := UI.FlowTrigSense
		//Begin loop
		for !Exit {
			//check if trigger is true
			if sensors.FIns.ReadFlow() >= FTrigger { //need to mkae sure of unit convention Lpm or mL
				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.TiMax*1000) * time.Millisecond); {
					valves.InProp.IncrementValve(PressurePID.Update(float64(sensors.PIns.ReadPressure())))
				}
				//Close main valve MIns
				valves.InProp.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled
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
	}

}
