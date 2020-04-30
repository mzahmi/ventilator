package modeselect

import (
	"time"

	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
)

// PSV one of the main 5 modes of the ventilator. The Triggering type is chosen by the operator.
// 	Triggering:	Pressure/Flow
// 	Cycling: 	Flow
// 	Control: 	Pressure
func PSV(UI *UserInput) {
	UpdateValues(UI)
	//Initialize  Sensors at inhalation side
	PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
	FIns := sensors.Flow{Name: "SNS_F_INS", ID: 1, AdcID: 1, Rate: 0}
	//Initialize  Sensors at exhalation side
	PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 2, AdcID: 1, MMH2O: 0} //expratory pressure sensor
	//FExp := sensors.Flow{Name: "SNS_F_EXP", ID: 3, AdcID: 1, Rate: 0}
	//Initialize valves
	MIns := valves.PropValve{Name: "A_PSV_INS", DacChan: 0, DacID: 1, Percent: 0} //normally closed
	MExp := valves.PropValve{Name: "A_PSV_EXP", DacChan: 1, DacID: 1, Percent: 0} //normally open
	//Initialize PID controller
	//initiate a PID controller based on the PeakFlow
	PressurePID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//Check trigger type
	switch UI.PatientTriggerType {
	case "Pressure":
		//Calculate trigger threshhold with PEEP and sensitivity
		PTrigger := UI.PEEP + UI.PressureTrigSense

		counter := 1
		//Begin loop
		for !Exit {
			FlowMonitor := []float32{FIns.ReadFlow()}
			//check if trigger is true
			if PIns.ReadPressure() <= PTrigger {
				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.TiMax*1000) * time.Millisecond); {
					MIns.IncrementValve(PressurePID.Update(float64(PIns.ReadPressure())))
					FlowMonitor = append(FlowMonitor, FIns.ReadFlow())
					if FlowMonitor[0] < FlowMonitor[counter] {
						FlowMonitor[0] = FlowMonitor[counter] // saves maximum value @ index 0
					} else if (FlowMonitor[0]*UI.FlowCyclePercent)/100 >= FIns.ReadFlow() {
						counter = 1
						FlowMonitor = nil
						break
					}
					counter++
				}
				//Close main valve MIns
				MIns.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled by flow sensor FExp
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
					if PExp.ReadPressure() <= UI.PEEP {
						break
					}
					MExp.IncrementValve(100)
				}
				//Close main valve MExp
				MExp.IncrementValve(0) // closes the valve
			}
		}
	case "Flow":
		//Calculate trigger threshhold with flow trig sensitivity
		FTrigger := UI.FlowTrigSense
		//Begin loop
		for !Exit {
			//check if trigger is true
			if FIns.ReadFlow() >= FTrigger {
				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
					MIns.IncrementValve(PressurePID.Update(float64(PIns.ReadPressure())))
				}
				//Close main valve MIns
				MIns.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled by flow sensor FExp
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
					if PExp.ReadPressure() <= UI.PEEP {
						break
					}
					MExp.IncrementValve(100)
				}
				//Close main valve MExp
				MExp.IncrementValve(0) // closes the valve
			}
		}
	}

}
