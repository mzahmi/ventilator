package modeselect

import (
	"fmt"
	"sync"

	"github.com/mzahmi/ventilator/control/alarms"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/params"
)

// params.UserInput is a custome type struct that contains the global
// variables input by the user or operator

// Exit is a global var used as a switch for github.com/mzahmi/ventilatorilation on or off
var Exit bool

// UpdateValues populates a a struct which is recieved by the GUI
func UpdateValues(UI *params.UserInput) {
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
	UI.PEEP = 10 * UI.PEEP                     // conversion from cmH2O to mmH2O
	UI.MinuteVolume = UI.TidalVolume * UI.Rate // calculation of minute volume MV = VT * BPM
}

// ModeSelection reads input from the GUI to select the required Mode from the user input struct
func ModeSelection(UI *params.UserInput, s chan sensors.SensorsReading, wg *sync.WaitGroup, readStatus chan string) {
	UpdateValues(UI) // calculates missing values
	switch UI.Mode {
	case "Pressure Control":
		fmt.Println("Pressure Control Mode selected")
		PressureAC(UI, s, wg, readStatus)
	case "Pressure A/C":
		fmt.Println("Pressure Assisted Control Mode selected")
	case "PSV":
		fmt.Println("Pressure Support Control Mode selected")
	case "V-SMIV":
		fmt.Println("Volume SIMV Mode selected")
	case "P-SIMV":
		fmt.Println("Pressure SIMV Mode selected")
	default:
		fmt.Println("No github.com/mzahmi/ventilatorilator Mode selected")
		return
	}
}

/*CheckAlarms ...*/
func CheckAlarms(UI *params.UserInput) error {
	errPIP := alarms.AirwayPressureAlarms(UI.UpperLimitPIP, UI.LowerLimitPIP)
	errVT := alarms.TidalVolumeAlarms(UI.UpperLimitVT, UI.LowerLimitVt)
	errMV := alarms.ExpiratoryMinuteVolumeAlarms(UI.UpperLimitMV, UI.LowerLimitMV)
	errRR := alarms.RespiratoryRateAlarms(UI.UpperLimitRR, UI.LowerLimitRR)

	if errPIP != nil {
		return errPIP
	} else if errVT != nil {
		return errVT
	} else if errMV != nil {
		return errMV
	} else if errRR != nil {
		return errRR
	} else {
		return nil
	}

}
