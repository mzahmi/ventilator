package modeselect

import (
	"sync"

	// "github.com/mzahmi/ventilator/control/alarms"
	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/logger"
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
func ModeSelection(UI *params.UserInput, s *sensors.SensorsReading, client *redis.Client, mux *sync.Mutex, logStruct *logger.Logging) {
	UpdateValues(UI) // calculates missing values
	//fmt.Println(UI.Mode)
	switch UI.Mode {
	case "Volume A/C":
		//logger.Println("Pressure Control Mode selected")
		logStruct.Event("Volume Assisted Control Mode selected")
		VolumeAC(UI, s, client, mux, logStruct)
	case "Pressure A/C":
		// logger.Println("Pressure Assisted Control Mode selected")
		logStruct.Event("Pressure Assisted Control Mode selected")
		PressureAC(UI, s, client, mux, logStruct)
	case "Pressure Support (PSV)":
		// logger.Println("Pressure Support Control Mode selected")
		logStruct.Event("Pressure Support Control Mode selected")
		PSV(UI, s, client, mux, logStruct)
	case "Volume SMIV":
		// logger.Println("Volume SIMV Mode selected")
		logStruct.Event("Volume SIMV Mode selected")
	case "Pressure SIMV":
		// logger.Println("Pressure SIMV Mode selected")
		logStruct.Event("Pressure SIMV Mode selected")
	default:
		// logger.Println("Incorrect Mode selected")
		logStruct.Event("Incorrect Mode selected")
		return
	}
}

// prints out the checked error err
func check(err error, logStruct *logger.Logging) {
	if err != nil {
		logStruct.Err(err)
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
