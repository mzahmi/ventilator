package satur

import (
	"fmt"
	"time"

	"github.com/mzahmi/ventilator/control/ioexp"
	"github.com/mzahmi/ventilator/control/rpigpio"

	// "github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
)

// UserInput is a custome type struct that contains the global
// variables input by the user or operator
type UserInput struct {
	// Mode                string
	// BreathType          string
	TidalVolume float32 // ml
	Rate        float32 // BPM
	Ti          float32 // inhalation time
	// TiMax               float32 // for PSV mode backup time control
	// Te                  float32 // exhalation time
	PeakFlow           float32
	IR                 float32 // inhalation ratio part
	ER                 float32 // exhalation ratio part
	PEEP               float32 // 5-20 mmH2O
	FiO2               float32 // 21% - 100%
	PatientTriggerType int
	PressureTrigSense  float32 // -0.5 to 02 mmH2O
	// FlowTrigSense       float32 // 0.5 to 5 Lpm
	// FlowCyclePercent    float32 // for flow cycling ranges from 0 to 100%
	// PressureSupport     float32 // needs to be defined
	// InspiratoryPressure float32
	// PressureControl     float32
}

// Exit is a global var used as a switch for github.com/mzahmi/ventilatorilation on or off
var Exit bool

// UpdateValues populates a a struct which is recieved by the GUI
// func UpdateValues(UI *UserInput) {
// 	BCT := 60 / UI.Rate
// 	if UI.Ti != 0 {
// 		UI.Te = BCT - UI.Ti
// 		UI.PeakFlow = (60 * UI.TidalVolume) / (UI.Ti * 1000)
// 	} else if UI.IR != 0 {
// 		UI.Ti = UI.IR / (UI.IR + UI.ER)
// 		UI.Te = BCT - UI.Ti
// 		UI.PeakFlow = (60 * UI.TidalVolume) / (UI.Ti * 1000)
// 	} else if UI.PeakFlow != 0 {
// 		UI.Ti = (60 * UI.TidalVolume) / (UI.PeakFlow * 1000)
// 		UI.Te = BCT - UI.Ti
// 	}
// }

// ModeSelection reads input from the GUI to select the
// required Mode using the UserInput.Mode
// func ModeSelection(UI *UserInput) {
// 	switch UI.Mode {
// 	case "Volume A/C":
// 		fmt.Println("Volume Assisted Control Mode selected")
// 		//VolumeAC(UI)
// 	case "Pressure A/C":
// 		fmt.Println("Pressure Assisted Control Mode selected")
// 	case "PSV":
// 		fmt.Println("Pressure Support Control Mode selected")
// 	case "V-SMIV":
// 		fmt.Println("Volume SIMV Mode selected")
// 	case "P-SIMV":
// 		fmt.Println("Pressure SIMV Mode selected")
// 	default:
// 		fmt.Println("No github.com/mzahmi/ventilatorilator Mode selected")
// 		return
// 	}
// }

// // VolumeAC is one of the main 5 modes of the github.com/mzahmi/ventilatorilator
// func VolumeAC(UI *UserInput) {
// 	switch UI.BreathType {
// 	case "Control":
// 		//VolumeControl(UI)
// 	case "Assist":
// 		//VolumeAssist(UI)
// 	default:
// 		fmt.Println("Enter valid breath type")
// 	}
// }

// // VolumeControl ...
// func VolumeControl(UI *UserInput) {
// 	//calculate Te from UI.Ti and BCT
// 	UpdateValues(UI)

// 	//initiate a PID controller based on the PeakFlow
// 	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values
// 	//Initialize  Sensors at inhalation side
// 	//PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
// 	FIns := sensors.Flow{Name: "SNS_F_INS", ID: 1, AdcID: 1, Rate: 0}
// 	//Initialize  Sensors at exhalation side
// 	PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 2, AdcID: 1, MMH2O: 0} //expratory pressure sensor
// 	//FExp := sensors.Flow{Name: "SNS_F_EXP", ID: 3, AdcID: 1, Rate: 0}
// 	//Initialize valves
// 	MIns := valves.PropValve{Name: "A_PSV_INS", DacChan: 0, DacID: 1, Percent: 0} //normally closed
// 	MExp := valves.PropValve{Name: "A_PSV_EXP", DacChan: 1, DacID: 1, Percent: 0} //normally open

// 	//control loop
// 	for !Exit {
// 		FlowPID.setpoint = float64(UI.PeakFlow) // Sets the PID setpoint
// 		//Open main valve MIns controlled by flow sensor PIns
// 		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
// 			MIns.IncrementValve(FlowPID.Update(float64(FIns.ReadFlow())))
// 		}
// 		//Close main valve MIns
// 		MIns.IncrementValve(0) // closes the valve
// 		//Open main valve MExp controlled by flow sensor PExp
// 		for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
// 			if PExp.ReadPressure() <= UI.PEEP {
// 				break
// 			}
// 			MExp.IncrementValve(100)
// 		}
// 		//Close main valve MExp
// 		MExp.IncrementValve(0) // closes the valve
// 	}

// }

// //VolumeAssist ...
// func VolumeAssist(UI *UserInput) {
// 	UpdateValues(UI)
// 	//Initialize  Sensors at inhalation side
// 	PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
// 	FIns := sensors.Flow{Name: "SNS_F_INS", ID: 1, AdcID: 1, Rate: 0}
// 	//Initialize  Sensors at exhalation side
// 	//PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 2, AdcID: 1, MMH2O: 0} //expratory pressure sensor
// 	FExp := sensors.Flow{Name: "SNS_F_EXP", ID: 3, AdcID: 1, Rate: 0}
// 	//Initialize valves
// 	MIns := valves.PropValve{Name: "A_PSV_INS", DacChan: 0, DacID: 1, Percent: 0} //normally closed
// 	MExp := valves.PropValve{Name: "A_PSV_EXP", DacChan: 1, DacID: 1, Percent: 0} //normally open
// 	//Initialize PID controller
// 	//initiate a PID controller based on the PeakFlow
// 	FlowPID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

// 	//Check trigger type
// 	switch UI.PatientTriggerType {
// 	case "Pressure":
// 		//Calculate trigger threshhold with PEEP and sensitivity
// 		PTrigger := UI.PEEP + UI.PressureTrigSense
// 		//Begin loop
// 		for !Exit {
// 			//check if trigger is true
// 			if PIns.ReadPressure() <= PTrigger {
// 				FlowPID.setpoint = float64(UI.PeakFlow) // Sets PID setpoint to PeakFlow
// 				//Open main valve MIns controlled by flow sensor FIns
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
// 					MIns.IncrementValve(FlowPID.Update(float64(FIns.ReadFlow())))
// 				}
// 				//Close main valve MIns
// 				MIns.IncrementValve(0) // closes the valve

// 				//Open main valve MExp controlled by flow sensor FExp
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te) * time.Second); {
// 					MExp.IncrementValve(FlowPID.Update(float64(FExp.ReadFlow())))
// 				}
// 				//Close main valve MExp
// 				MExp.IncrementValve(0) // closes the valve
// 			}
// 		}
// 	case "Flow":
// 		//Calculate trigger threshhold with PEEP and sensitivity
// 		FTrigger := UI.FlowTrigSense
// 		//Begin loop
// 		for !Exit {
// 			//check if trigger is true
// 			if PIns.ReadPressure() >= FTrigger {
// 				FlowPID.setpoint = float64(UI.PeakFlow) // Sets PID setpoint to PeakFlow
// 				//Open main valve MIns controlled by flow sensor FIns
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti) * time.Second); {
// 					MIns.IncrementValve(FlowPID.Update(float64(FIns.ReadFlow())))
// 				}
// 				//Close main valve MIns
// 				MIns.IncrementValve(0) // closes the valve

// 				//Open main valve MExp controlled by flow sensor FExp
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te) * time.Second); {
// 					MExp.IncrementValve(FlowPID.Update(float64(FExp.ReadFlow())))
// 				}
// 				//Close main valve MExp
// 				MExp.IncrementValve(0) // closes the valve
// 			}
// 		}

// 	}
// }

// // PressureAC ...
// func PressureAC(UI *UserInput) {
// 	switch UI.BreathType {
// 	case "Control":
// 		PressureControl(UI)
// 	case "Assist":
// 		PressureAssist(UI)
// 	default:
// 		fmt.Println("Enter valid breath type")
// 	}
// }

// PressureControl ...
func PressureControl(UI *UserInput) {
	fmt.Println("entered function")

	//calculate Te from UI.Ti and BCT
	// UpdateValues(UI)
	BCT := 60 / UI.Rate
	Te := BCT - UI.Ti
	UI.PeakFlow = (60 * UI.TidalVolume) / (UI.Ti * 1000)

	// Identify the main valves or solenoids by MIns and MExp
	Mv := valves.SolenValve{Name: "main", State: false, PinMask: ioexp.Solenoid0}        //normally closed
	MIns := valves.SolenValve{Name: "A_PSV_INS", State: false, PinMask: ioexp.Solenoid1} //normally closed
	MExp := valves.SolenValve{Name: "A_PSV_EXP", State: true, PinMask: ioexp.Solenoid2}  //normally open

	// Identify the flow sensors by PIns and PExp
	//PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
	//PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 1, AdcID: 1, MMH2O: 0} //expratory pressure sensor

	fmt.Println("beep called")
	err := rpigpio.BeepOn()
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1000 * time.Millisecond)
	err = rpigpio.BeepOff()
	if err != nil {
		fmt.Println(err)
		return
	}

	//control loop
	for {
		fmt.Println("entered super for loop")

		//Open main valve MIns controlled by flow sensor PIns
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
			fmt.Println("INHALE")
			err := ioexp.WritePin(ioexp.GreenLed, true)
			if err != nil {
				fmt.Println(err)
				return
			}

			Mv.SolenCmd("Open")
			MIns.SolenCmd("Open")
		}
		//Close main valve MIns
		MIns.SolenCmd("Close") // closes the valve
		Mv.SolenCmd("Close")

		err = ioexp.WritePin(ioexp.GreenLed, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Open main valve MExp controlled by flow sensor PExp
		for start := time.Now(); time.Since(start) < (time.Duration(Te*1000) * time.Millisecond); {
			err = ioexp.WritePin(ioexp.RedLed, true)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("EXHALE")
			MExp.SolenCmd("Open")
		}
		err = ioexp.WritePin(ioexp.RedLed, false)
		if err != nil {
			fmt.Println(err)
			return
		}
		//Close main valve MExp
		MExp.SolenCmd("Close") // closes the valve
	}

}

// // PressureAssist ...
// func PressureAssist(UI *UserInput) {

// 	UpdateValues(UI)
// 	//Initialize  Sensors at inhalation side
// 	PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
// 	FIns := sensors.Flow{Name: "SNS_F_INS", ID: 1, AdcID: 1, Rate: 0}
// 	//Initialize  Sensors at exhalation side
// 	PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 2, AdcID: 1, MMH2O: 0} //expratory pressure sensor
// 	//FExp := sensors.Flow{Name: "SNS_F_EXP", ID: 3, AdcID: 1, Rate: 0}
// 	//Initialize valves
// 	MIns := valves.PropValve{Name: "A_PSV_INS", DacChan: 0, DacID: 1, Percent: 0} //normally closed
// 	MExp := valves.PropValve{Name: "A_PSV_EXP", DacChan: 1, DacID: 1, Percent: 0} //normally open
// 	//Initialize PID controller
// 	//initiate a PID controller based on the PeakFlow
// 	PressurePID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

// 	//Check trigger type
// 	switch UI.PatientTriggerType {
// 	case "Pressure":
// 		//Calculate trigger threshhold with PEEP and sensitivity
// 		PTrigger := UI.PEEP + UI.PressureTrigSense
// 		//Begin loop
// 		for !Exit {
// 			//check if trigger is true
// 			if PIns.ReadPressure() <= PTrigger {
// 				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
// 				//Open main valve MIns controlled by flow sensor FIns
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
// 					MIns.IncrementValve(PressurePID.Update(float64(PIns.ReadPressure())))
// 				}
// 				//Close main valve MIns
// 				MIns.IncrementValve(0) // closes the valve

// 				//Open main valve MExp controlled by flow sensor FExp
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
// 					if PExp.ReadPressure() <= UI.PEEP {
// 						break
// 					}
// 					MExp.IncrementValve(100)
// 				}
// 				//Close main valve MExp
// 				MExp.IncrementValve(0) // closes the valve
// 			}
// 		}
// 	case "Flow":
// 		//Calculate trigger threshhold with flow trig sensitivity
// 		FTrigger := UI.FlowTrigSense
// 		//Begin loop
// 		for !Exit {
// 			//check if trigger is true
// 			if FIns.ReadFlow() >= FTrigger {
// 				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
// 				//Open main valve MIns controlled by flow sensor FIns
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
// 					MIns.IncrementValve(PressurePID.Update(float64(PIns.ReadPressure())))
// 				}
// 				//Close main valve MIns
// 				MIns.IncrementValve(0) // closes the valve

// 				//Open main valve MExp controlled by flow sensor FExp
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
// 					if PExp.ReadPressure() <= UI.PEEP {
// 						break
// 					}
// 					MExp.IncrementValve(100)
// 				}
// 				//Close main valve MExp
// 				MExp.IncrementValve(0) // closes the valve
// 			}
// 		}
// 	}
// }

// // PSV Pressure support github.com/mzahmi/ventilatorilation mode
// func PSV(UI *UserInput) {
// 	UpdateValues(UI)
// 	//Initialize  Sensors at inhalation side
// 	PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0} //insparatory pressure sensor
// 	FIns := sensors.Flow{Name: "SNS_F_INS", ID: 1, AdcID: 1, Rate: 0}
// 	//Initialize  Sensors at exhalation side
// 	PExp := sensors.Pressure{Name: "SNS_P_EXP", ID: 2, AdcID: 1, MMH2O: 0} //expratory pressure sensor
// 	//FExp := sensors.Flow{Name: "SNS_F_EXP", ID: 3, AdcID: 1, Rate: 0}
// 	//Initialize valves
// 	MIns := valves.PropValve{Name: "A_PSV_INS", DacChan: 0, DacID: 1, Percent: 0} //normally closed
// 	MExp := valves.PropValve{Name: "A_PSV_EXP", DacChan: 1, DacID: 1, Percent: 0} //normally open
// 	//Initialize PID controller
// 	//initiate a PID controller based on the PeakFlow
// 	PressurePID := NewPIDController(0.5, 0.5, 0.5) // takes in P, I, and D values

// 	//Check trigger type
// 	switch UI.PatientTriggerType {
// 	case "Pressure":
// 		//Calculate trigger threshhold with PEEP and sensitivity
// 		PTrigger := UI.PEEP + UI.PressureTrigSense

// 		counter := 1
// 		//Begin loop
// 		for !Exit {
// 			FlowMonitor := []float32{FIns.ReadFlow()}
// 			//check if trigger is true
// 			if PIns.ReadPressure() <= PTrigger {
// 				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
// 				//Open main valve MIns controlled by flow sensor FIns
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.TiMax*1000) * time.Millisecond); {
// 					MIns.IncrementValve(PressurePID.Update(float64(PIns.ReadPressure())))
// 					FlowMonitor = append(FlowMonitor, FIns.ReadFlow())
// 					if FlowMonitor[0] < FlowMonitor[counter] {
// 						FlowMonitor[0] = FlowMonitor[counter] // saves maximum value @ index 0
// 					} else if (FlowMonitor[0]*UI.FlowCyclePercent)/100 >= FIns.ReadFlow() {
// 						counter = 1
// 						FlowMonitor = nil
// 						break
// 					}
// 					counter++
// 				}
// 				//Close main valve MIns
// 				MIns.IncrementValve(0) // closes the valve

// 				//Open main valve MExp controlled by flow sensor FExp
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
// 					if PExp.ReadPressure() <= UI.PEEP {
// 						break
// 					}
// 					MExp.IncrementValve(100)
// 				}
// 				//Close main valve MExp
// 				MExp.IncrementValve(0) // closes the valve
// 			}
// 		}
// 	case "Flow":
// 		//Calculate trigger threshhold with flow trig sensitivity
// 		FTrigger := UI.FlowTrigSense
// 		//Begin loop
// 		for !Exit {
// 			//check if trigger is true
// 			if FIns.ReadFlow() >= FTrigger {
// 				PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure
// 				//Open main valve MIns controlled by flow sensor FIns
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
// 					MIns.IncrementValve(PressurePID.Update(float64(PIns.ReadPressure())))
// 				}
// 				//Close main valve MIns
// 				MIns.IncrementValve(0) // closes the valve

// 				//Open main valve MExp controlled by flow sensor FExp
// 				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
// 					if PExp.ReadPressure() <= UI.PEEP {
// 						break
// 					}
// 					MExp.IncrementValve(100)
// 				}
// 				//Close main valve MExp
// 				MExp.IncrementValve(0) // closes the valve
// 			}
// 		}
// 	}

// }
