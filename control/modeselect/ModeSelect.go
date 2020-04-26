package modeselect

import (
	"fmt"
	"time"

	"vent/control/sensors"
	"vent/control/valves"
)

// UserInput is a custome type struct that contains the global
// variables input by the user or operator
type UserInput struct {
	Mode                string
	BreathType          string
	PatientTriggerType  string
	TidalVolume         float32 // ml
	Rate                float32 // BPM
	Ti                  float32 // inhalation time
	IR                  float32 // inhalation ratio part
	ER                  float32 // exhalation ratio part
	PeakFlow            float64
	PEEP                float32 // 5-20 mmH2O
	FiO2                float32 // 21% - 100%
	PressureTrigSense   float32 // -0.5 to 02 mmH2O
	FlowTrigSense       float32 // 0.5 to 5 Lpm
	PressureSupport     float32 // needs to be defined
	InspiratoryPressure float32
	PressureControl     float32
	Exit                bool
}

// ModeSelection reads input from the GUI to select the
// required Mode
func ModeSelection(UI *UserInput) {
	switch UI.Mode {
	case "Volume A/C":
		fmt.Println("Volume Assisted Control Mode selected")
		VolumeAC(UI)
	case "Pressure A/C":
		fmt.Println("Pressure Assisted Control Mode selected")
	case "PSV":
		fmt.Println("Pressure Support Control Mode selected")
	case "V-SMIV":
		fmt.Println("Volume SIMV Mode selected")
	case "P-SIMV":
		fmt.Println("Pressure SIMV Mode selected")
	default:
		fmt.Println("No Ventilator Mode selected")
		return
	}
}

// VolumeAC ...
func VolumeAC(UI *UserInput) {
	//Check breath type and run
	switch UI.BreathType {
	case "control":
		VolumeControl(UI)
	case "assist":
		VolumeAssist(UI)
	default:
		fmt.Println("Enter valid breath type")
	}
}

// VolumeControl ...
func VolumeControl(UI *UserInput) {
	//calculate Te from UI.Ti and BCT
	BCT := 60 / UI.Rate
	Te := BCT - UI.Ti

	//initiate a PID controller based on the PeakFlow
	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values
	FlowPID = FlowPID.Set(UI.PeakFlow)         // Sets the PID setpoint

	// Identify the main valves or solenoids by MIns and MExp
	MIns := valves.PropValve{ID: "A_PSV_INS", Address: "GPIO01", Percent: 0}   //normally closed
	MExp := valves.PropValve{ID: "A_PSV_EXP", Address: "GPIO02", Percent: 100} //normally open

	// Identify the flow sensors by FIns and FExp
	FIns := sensors.Flow{ID: "SNS_F_INS", Address: "GPIO03", Rate: 0}
	FExp := sensors.Flow{ID: "SNS_F_EXP", Address: "GPIO04", Rate: 0}

	//control loop
	for UI.Exit == false {
		//Open main valve MIns controlled by flow sensor FIns
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
			MIns.IncrementValve(FlowPID.Update(FIns.ReadFlow()))
		}
		//Close main valve MIns
		MIns.IncrementValve(0) // closes the valve
		//Open main valve MExp controlled by flow sensor FExp
		for start := time.Now(); time.Since(start) < (time.Duration(Te) * time.Second); {
			MExp.IncrementValve(FlowPID.Update(FExp.ReadFlow()))
		}
		//Close main valve MExp
		MExp.IncrementValve(0) // closes the valve
	}

}

// VolumeAssist ...
func VolumeAssist(UI *UserInput) {
	//Initialize  Sensors
	PIns := sensors.Pressure{ID: "Inhalation Pressure Sensor", Address: "GPIOX", MMH2O: 0}
	FIns := sensors.Flow{ID: "SNS_F_INS", Address: "GPIO03", Rate: 0}
	FExp := sensors.Flow{ID: "SNS_F_EXP", Address: "GPIO04", Rate: 0}
	//Initialize valves
	MIns := valves.PropValve{ID: "A_PSV_INS", Address: "GPIO01", Percent: 0}   //normally closed
	MExp := valves.PropValve{ID: "A_PSV_EXP", Address: "GPIO02", Percent: 100} //normally open
	//Initialize PID controller
	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//Check trigger type
	switch UI.PatientTriggerType {
	case "pressure":
		//Calculate trigger threshhold with PEEP and sensitivity
		PTrigger := UI.PEEP + UI.PressureTrigSense
		//Ti, IE, or Peak Flow Cycling
		if UI.Ti != 0 {
			//Calculate Te based on BCT and Ti
			Te := 60/UI.Rate - UI.Ti
			//Calulcate minimum flow rate(L/min) to deliver Tidal Volume(mL) given Ti(s)
			FlowRate := UI.TidalVolume * 60 / (UI.Ti * 1000)
			//Setpoint for PID
			FlowPID.setpoint = float64(FlowRate)
			//Begin loop
			for !UI.Exit {
				//check if trigger is true
				if PIns.ReadPressure() <= PTrigger {
					//Open main valve MIns controlled by flow sensor FIns
					for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
						MIns.IncrementValve(FlowPID.Update(float64(FIns.ReadFlow())))
					}
					//Close main valve MIns
					MIns.IncrementValve(0) // closes the valve
					//Open main valve MExp controlled by flow sensor FExp
					for start := time.Now(); time.Since(start) < (time.Duration(Te) * time.Second); {
						MExp.IncrementValve(FlowPID.Update(float64(FExp.ReadFlow())))
					}
					//Close main valve MExp
					MExp.IncrementValve(0) // closes the valve
				}
			}
		} else if UI.PeakFlow != 0 {

		} else if UI.IR != 0 {

		}

	}
}
