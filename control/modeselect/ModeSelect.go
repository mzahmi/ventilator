package modeselect

import (
	"fmt"
	"time"

	"github.com/mzahmi/ventilator/control/ioexp"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
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
	Te                  float32 // exhalation time
	IR                  float32 // inhalation ratio part
	ER                  float32 // exhalation ratio part
	PeakFlow            float32
	PEEP                float32 // 5-20 mmH2O
	FiO2                float32 // 21% - 100%
	PressureTrigSense   float32 // -0.5 to 02 mmH2O
	FlowTrigSense       float32 // 0.5 to 5 Lpm
	PressureSupport     float32 // needs to be defined
	InspiratoryPressure float32
	PressureControl     float32
}

// Exit is a global var used as a switch for ventilation on or off
var Exit bool

// UpdateValues ...
func UpdateValues(UI *UserInput) {
	BCT := 60 / UI.Rate
	if UI.Ti != 0 {
		UI.Te = BCT - UI.Ti
		UI.PeakFlow = (60 * UI.TidalVolume) / (UI.Ti * 1000)
	} else if UI.IR != 0 {
		UI.Ti = UI.IR / (UI.IR + UI.ER)
		UI.Te = BCT - UI.Ti
		UI.PeakFlow = (60 * UI.TidalVolume) / (UI.Ti * 1000)
	} else if UI.PeakFlow != 0 {
		UI.Ti = (60 * UI.TidalVolume) / (UI.PeakFlow * 1000)
		UI.Te = BCT - UI.Ti
	}
}

// ModeSelection reads input from the GUI to select the
// required Mode
func ModeSelection(UI *UserInput) {
	switch UI.Mode {
	case "Volume A/C":
		fmt.Println("Volume Assisted Control Mode selected")
		//VolumeAC(UI)
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
	switch UI.BreathType {
	case "Control":
		//VolumeControl(UI)
	case "Assist":
		//VolumeAssist(UI)
	default:
		fmt.Println("Enter valid breath type")
	}
}

// VolumeControl ...
func VolumeControl(UI *UserInput) {
	//calculate Te from UI.Ti and BCT
	UpdateValues(UI)

	//initiate a PID controller based on the PeakFlow
	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values
	//Initialize  Sensors at inhalation side
	//PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
	FIns := sensors.Flow{Name: "SNS_F_INS", ID: 1, AdcID: 1, Rate: 0}
	//Initialize  Sensors at exhalation side
	PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 2, AdcID: 1, MMH2O: 0} //expratory pressure sensor
	//FExp := sensors.Flow{Name: "SNS_F_EXP", ID: 3, AdcID: 1, Rate: 0}
	//Initialize valves
	MIns := valves.PropValve{Name: "A_PSV_INS", DacChan: 0, DacID: 1, Percent: 0} //normally closed
	MExp := valves.PropValve{Name: "A_PSV_EXP", DacChan: 1, DacID: 1, Percent: 0} //normally open

	//control loop
	for !Exit {
		FlowPID.setpoint = float64(UI.PeakFlow) // Sets the PID setpoint
		//Open main valve MIns controlled by flow sensor PIns
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
			MIns.IncrementValve(FlowPID.Update(float64(FIns.ReadFlow())))
		}
		//Close main valve MIns
		MIns.IncrementValve(0) // closes the valve
		//Open main valve MExp controlled by flow sensor PExp
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

//VolumeAssist ...
func VolumeAssist(UI *UserInput) {
	UpdateValues(UI)
	//Initialize  Sensors at inhalation side
	PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
	FIns := sensors.Flow{Name: "SNS_F_INS", ID: 1, AdcID: 1, Rate: 0}
	//Initialize  Sensors at exhalation side
	//PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 2, AdcID: 1, MMH2O: 0} //expratory pressure sensor
	FExp := sensors.Flow{Name: "SNS_F_EXP", ID: 3, AdcID: 1, Rate: 0}
	//Initialize valves
	MIns := valves.PropValve{Name: "A_PSV_INS", DacChan: 0, DacID: 1, Percent: 0} //normally closed
	MExp := valves.PropValve{Name: "A_PSV_EXP", DacChan: 1, DacID: 1, Percent: 0} //normally open
	//Initialize PID controller
	//initiate a PID controller based on the PeakFlow
	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

	//Check trigger type
	switch UI.PatientTriggerType {
	case "Pressure":
		//Calculate trigger threshhold with PEEP and sensitivity
		PTrigger := UI.PEEP + UI.PressureTrigSense
		//Begin loop
		for !Exit {
			//check if trigger is true
			if PIns.ReadPressure() <= PTrigger {
				FlowPID.setpoint = float64(UI.PeakFlow) // Sets PID setpoint to PeakFlow
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
					MIns.IncrementValve(FlowPID.Update(float64(FIns.ReadFlow())))
				}
				//Close main valve MIns
				MIns.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled by flow sensor FExp
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te) * time.Second); {
					MExp.IncrementValve(FlowPID.Update(float64(FExp.ReadFlow())))
				}
				//Close main valve MExp
				MExp.IncrementValve(0) // closes the valve
			}
		}
	case "Flow":
		//Calculate trigger threshhold with PEEP and sensitivity
		FTrigger := UI.FlowTrigSense
		//Begin loop
		for !Exit {
			//check if trigger is true
			if PIns.ReadPressure() >= FTrigger {
				FlowPID.setpoint = float64(UI.PeakFlow) // Sets PID setpoint to PeakFlow
				//Open main valve MIns controlled by flow sensor FIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
					MIns.IncrementValve(FlowPID.Update(float64(FIns.ReadFlow())))
				}
				//Close main valve MIns
				MIns.IncrementValve(0) // closes the valve

				//Open main valve MExp controlled by flow sensor FExp
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te) * time.Second); {
					MExp.IncrementValve(FlowPID.Update(float64(FExp.ReadFlow())))
				}
				//Close main valve MExp
				MExp.IncrementValve(0) // closes the valve
			}
		}

	}
}

// PressureAC ...
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

// PressureControl ...
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

// PressureAssist ...
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
