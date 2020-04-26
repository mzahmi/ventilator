package modeselect

import (
	"fmt"
	"github.com/mzahmi/ventilator/control/valves"
	"github.com/mzahmi/ventilator/control/sensors"
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
	FlowPID := NewPIDController(0.5, 0.5, 0.5)
	FlowPID = FlowPID.Set(UI.PeakFlow)
	F1 := valves.PropValve{1,"flow actuator",0}
	Fsense := sensors.Flow{1,1,1,1}
	for start:=time.Now(); time.Since(start) < (UI.Ti * time.Second) {
		F1.IncrementValve(FlowPID.Update(Fsense.ReadFlow()))

	}
	F1.IncrementValve(0) // closes the valve
	

}

// VolumeAssist ...
func VolumeAssist(UI *UserInput) {

}
