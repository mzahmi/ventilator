package modeselect

import (
	"fmt"
	"time"

	"github.com/mzahmi/ventilator/control/ioexp"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
)

//PressureACSettings is a struct that contains all the variables
//that will/can be used for Pressure AC mode
type PressureACSettings struct {
	BreathType          string
	PatientTriggerType  string
	Rate                float32 // BPM
	Ti                  float32 // inhalation time
	TiMax               float32 // for PSV mode backup time control
	Te                  float32 // exhalation time
	IR                  float32 // inhalation ratio part
	ER                  float32 // exhalation ratio part
	PEEP                float32 // 5-20 mmH2O
	FiO2                float32 // 21% - 100%
	PressureTrigSense   float32 // -0.5 to 02 mmH2O
	FlowTrigSense       float32 // 0.5 to 5 Lpm
	InspiratoryPressure float32
}

// PressureAC one of the main 5 modes of the ventilator
func PressureAC(UI *UserInput) {
	switch UI.BreathType {
	case "Control":
		PressureControl(UI)
	case "Assist":
		PressureAssist(UI)
	default:
		fmt.Println("Enter valid breath type")
	}
}

// PressureControl a submode of PressureAC.
// 	Triggering:	Time
// 	Cycling: 	Time
// 	Control: 	Pressure
func PressureControl(UI *UserInput) {
	//calculate Te from UI.Ti and BCT
	UpdateValues(UI)

	// Identify the main valves or solenoids by MIns and MExp
	MIns := valves.SolenValve{Name: "A_PSV_INS", State: false, PinMask: ioexp.Solenoid0} //normally closed
	MExp := valves.SolenValve{Name: "A_PSV_EXP", State: true, PinMask: ioexp.Solenoid1}  //normally open

	// Identify the flow sensors by PIns and PExp
	//PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
	//PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 1, AdcID: 1, MMH2O: 0} //expratory pressure sensor

	//control loop
	for !Exit {
		//Open main valve MIns controlled by flow sensor PIns
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
			MIns.SolenCmd("Open")
		}
		//Close main valve MIns
		MIns.SolenCmd("Close") // closes the valve
		//Open main valve MExp controlled by flow sensor PExp
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
			MExp.SolenCmd("Open")
		}
		//Close main valve MExp
		MExp.SolenCmd("Close") // closes the valve
	}

}

// PressureAssist a submode of PressureAC. The Triggering type is chosen by the operator.
// 	Triggering:	Pressure/Flow
// 	Cycling: 	Time
// 	Control: 	Pressure
func PressureAssist(UI *UserInput) {

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
		//Begin loop
		for !Exit {
			//check if trigger is true
			if PIns.ReadPressure() <= PTrigger {
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
