package modeselect

import (
	"fmt"
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
	TiMax               float32 // for PSV mode backup time control
	Te                  float32 // exhalation time
	IR                  float32 // inhalation ratio part IR:IE
	ER                  float32 // exhalation ratio part IR:IE
	PeakFlow            float32
	PEEP                float32 // 5-20 mmH2O
	FiO2                float32 // 21% - 100%
	PressureTrigSense   float32 // -0.5 to 02 mmH2O
	FlowTrigSense       float32 // 0.5 to 5 Lpm
	FlowCyclePercent    float32 // for flow cycling ranges from 0 to 100%
	PressureSupport     float32 // needs to be defined
	InspiratoryPressure float32 // Also known as P_control
	UpperLimitVT        float32 // upper limit of tidal volume
	LowerLimitVt        float32 // lower limit of tidal volume
	RiseTime            float32 // needs to be defined
	UpperLimitPIP       float32 // upper limit of airway peak prssure
	LowerLimitPIP       float32 // lower limit of airway peak pressure
}

// Exit is a global var used as a switch for ventilation on or off
var Exit bool

// UpdateValues populates a a struct which is recieved by the GUI
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
	UI.PEEP = 10 * UI.PEEP // conversion from cmH2O to mmH2O
}

// ModeSelection reads input from the GUI to select the required Mode from the user input struct
func ModeSelection(UI *UserInput) {
	UpdateValues(UI) // calculates missing values
	switch UI.Mode {
	case "Volume A/C":
		fmt.Println("Volume Assisted Control Mode selected")
		VolumeAC(UI)
	case "Pressure A/C":
		fmt.Println("Pressure Assisted Control Mode selected")
		PressureAC(UI)
	case "PSV":
		fmt.Println("Pressure Support Control Mode selected")
		PSV(UI)
	case "V-SMIV":
		fmt.Println("Volume SIMV Mode selected")
	case "P-SIMV":
		fmt.Println("Pressure SIMV Mode selected")
	default:
		fmt.Println("No Ventilator Mode selected")
		return
	}
}
