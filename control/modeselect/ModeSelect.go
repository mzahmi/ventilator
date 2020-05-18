package modeselect

import (
	"fmt"
	"log"
	"sync"

	// "github.com/mzahmi/ventilator/control/alarms"
	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/params"
)

// params.UserInput is a custome type struct that contains the global

// UpdateValues populates a a struct which is recieved by the GUI
func UpdateValues(UI *params.UserInput) {
	BCT := 60 / UI.Rate
	UI.Ti = BCT / (1 + UI.ER) // UI.IR = 1 in this case until resolved
	UI.Te = BCT - UI.Ti
	UI.MinuteVolume = UI.TidalVolume * UI.Rate // calculation of minute volume MV = VT * BPM
}

// ModeSelection reads input from the GUI to select the required Mode from the user input struct
func ModeSelection(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logger *log.Logger) {
	UpdateValues(UI) // calculates missing values
	//fmt.Println(UI.Mode)
	switch UI.Mode {
	case "Volume A/C":
		fmt.Println("Pressure Control Mode selected")
		VolumeAC(UI, s, client, mux, logger)
	case "Pressure A/C":
		logger.Println("Pressure Assisted Control Mode selected")
		PressureAC(UI, s, client, mux, logger)
	case "Pressure Support (PSV)":
		fmt.Println("Pressure Support Control Mode selected")
		PSV(UI, s, client, mux, logger)
	case "Volume SMIV":
		fmt.Println("Volume SIMV Mode selected")
	case "Pressure SIMV":
		fmt.Println("Pressure SIMV Mode selected")
	default:
		fmt.Println("No github.com/mzahmi/ventilator Mode selected")
		return
	}
}

// prints out the checked error err
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// /*CheckAlarms ...*/
// func CheckAlarms(UI *params.UserInput, s chan sensors.SensorsReading) error {
// 	errPIP := alarms.AirwayPressureAlarms(s,UI.UpperLimitPIP, UI.LowerLimitPIP)
// 	errVT := alarms.TidalVolumeAlarms(UI.UpperLimitVT, UI.LowerLimitVT)
// 	errMV := alarms.ExpiratoryMinuteVolumeAlarms(UI.UpperLimitMV, UI.LowerLimitMV)
// 	errRR := alarms.RespiratoryRateAlarms(UI.UpperLimitRR, UI.LowerLimitRR)

// 	if errPIP != nil {
// 		return errPIP
// 	} else if errVT != nil {
// 		return errVT
// 	} else if errMV != nil {
// 		return errMV
// 	} else if errRR != nil {
// 		return errRR
// 	} else {
// 		return nil
// 	}

// }
