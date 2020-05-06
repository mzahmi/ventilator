package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mzahmi/ventilator/control/modeselect"
	"github.com/mzahmi/ventilator/control/sensors"

	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

var UI = modeselect.UserInput{
	Mode:                "NA",
	BreathType:          "NA",
	PatientTriggerType:  "NA",
	TidalVolume:         0,
	Rate:                0,
	Ti:                  0,
	TiMax:               0,
	Te:                  0,
	IR:                  0,
	ER:                  0,
	PeakFlow:            0,
	PEEP:                0,
	FiO2:                0,
	PressureTrigSense:   0,
	FlowTrigSense:       0,
	FlowCyclePercent:    0,
	PressureSupport:     0,
	InspiratoryPressure: 0,
	UpperLimitVT:        0,
	LowerLimitVt:        0,
	RiseTime:            0,
	UpperLimitPIP:       0,
	LowerLimitPIP:       0,
	MinuteVolume:        0,
	UpperLimitMV:        0,
	LowerLimitMV:        0,
	UpperLimitRR:        0,
	LowerLimitRR:        0,
}

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
