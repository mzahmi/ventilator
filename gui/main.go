/*
singal : from go to qml
slot : from qml to go
*/

package main

import (
	"fmt"
	// "time"
	"github.com/mzahmi/ventilator/demo/satur"
	"github.com/therecipe/qt/gui"
)

func main() {

	// BUTTON connection:
	// Connects all input variables from qml to go
	qmlBridge.ConnectSendToGo(func(tidalVolume float32, rate float32, ti float32, peakFlow float32, ir float32, er float32, peep float32, fio2 float32, triggerType int, trigSense float32) {
		fmt.Println(tidalVolume, rate, ti, peakFlow, ir, er, peep, fio2, triggerType, trigSense)
		qmlBridge.SendInfo("Success")

		// make a struct from the user input
		ui := satur.UserInput{
			TidalVolume:        tidalVolume,
			Rate:               rate,
			Ti:                 ti,
			PeakFlow:           peakFlow,
			IR:                 ir,
			ER:                 er,
			PEEP:               peep,
			PatientTriggerType: triggerType,
			PressureTrigSense:  trigSense,
		}

		go satur.PressureControl(&ui)

	})

	// BUTTON connection:
	// Free Button
	qmlBridge.ConnectSendModeSelected(func(mode string) {
		fmt.Println(fmt.Println("Mode selected is ", mode))
	})

	// SEND UI
	// sends information to QML to display
	// counter := 0
	// read from a file every second
	// go func() {
	// 	i := 0
	// 	for i <= 100 {
	// 		qmlBridge.SendToQml(ReadFromFile())
	// 		time.Sleep(time.Second)
	// 		counter++
	// 	}
	// }()

	// WRAPPER
	// keep program running until exit window
	InitView().Show()
	gui.QGuiApplication_Exec()
}
