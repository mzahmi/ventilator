package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/mzahmi/ventilator/control/modeselect"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	InitializeCharts()
	app := widgets.NewQApplication(len(os.Args), os.Args)
	engine := qml.NewQQmlApplicationEngine(nil)
	var qmlBridge = NewQmlBridge(nil)

	engine.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	engine.Load(core.NewQUrl3("qrc:/qml/MainQt.qml", 0))

	qmlBridge.ConnectSendToGo(func(data string) string {
		fmt.Println("go:", data)
		return "hello from go"
	})

	var wg sync.WaitGroup

	go func() {
		for range time.NewTicker(time.Millisecond * 50).C {
			pressurevalue := int(sensors.PIns.ReadPressure())
			qmlBridge.SendToQml(pressurevalue)
		}
	}()

	//read sensors check alarms and send data to gui
	go func() {
		for {
			//Read sensors
			pip := int(sensors.PIns.ReadPressure())
			vt := int(sensors.FIns.ReadFlow())
			rate := int(UI.Rate)
			peep := int(sensors.PExp.ReadPressure())
			fio2 := int(UI.FiO2)
			mode := UI.Mode
			//send values to GUI
			qmlBridge.SendMonitor(pip, vt, rate, peep, fio2, mode)
		}

	}()

	//Receive mode settings and run mode
	go func() {
		qmlBridge.ConnectSendPAC(func(ie int, pip int, bpm int, pmax int, peep int, fio2 int) {
			RecieveModeSettings(ie, pip, bpm, pmax, peep, fio2)
		})
		modeselect.ModeSelection(&UI)
	}()

	// Enable the CLI
	/*
		wg.Add(1)
		ch := make(chan UserInput)
		go cli(&wg, ch)
		wg.Wait()
	*/
	app.Exec()
}

func RecieveModeSettings(ie int, pip int, bpm int, pmax int, peep int, fio2 int) {
	switch ie {
	case 0:
		UI.IR = 1
		UI.ER = 2
	case 1:
		UI.IR = 1
		UI.ER = 3
	default:
		fmt.Println("invalid IE ratio")
	}
	UI.UpperLimitPIP = float32(pip)
	UI.Rate = float32(bpm)
	UI.InspiratoryPressure = float32(pmax)
	UI.PEEP = float32(peep)
	UI.FiO2 = float32(fio2)
}

func InitializeCharts() { _ = charts.QChart{} }

type QmlBridge struct {
	core.QObject

	_ func(data int)                                                   `signal:"sendToQml"`
	_ func(pip int, vt int, rate int, peep int, fio2 int, mode string) `signal:"sendMonitor"`
	_ func(data string) string                                         `slot:"sendToGo"`
	_ func(ie int, pip int, bpm int, pmax int, peep int, fio2 int)     `slot:"sendPAC"`
}
