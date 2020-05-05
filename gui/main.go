package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	// "github.com/mzahmi/ventilator/control/modeselect"
	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

func RecieveModeSettings(ie int, pip int, bpm int, pmax int, peep int, fio2 int) {
	fmt.Println(ie, pip, bpm, pmax, peep, fio2)
}

func InitializeCharts() { _ = charts.QChart{} }

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

	qmlBridge.ConnectSendPAC(func(ie int, pip int, bpm int, pmax int, peep int, fio2 int) {
		RecieveModeSettings(ie, pip, bpm, pmax, peep, fio2)
	})

	go func() {
		for range time.NewTicker(time.Millisecond * 50).C {
			randnumber := rand.Intn(30)
			qmlBridge.SendToQml(randnumber)
		}
	}()

	go func() {
		for range time.NewTicker(time.Second * 1).C {

			pip := rand.Intn(30)
			vt := rand.Intn(30)
			rate := rand.Intn(30)
			peep := rand.Intn(30)
			fio2 := rand.Intn(30)
			mode := "PAC"
			qmlBridge.SendMonitor(pip, vt, rate, peep, fio2, mode)
		}
	}()

	// read sensors check alarms and send data to gui
	// go readingAndDisplay()
	// {
	// 	//sensors read all
	// 	//check alarms
	// 	//send/share to GUI
	// }

	// // run the required ventilation mode
	// go ventilationControl()

	// // run command line interface
	// go cli()

	app.Exec()
}

type QmlBridge struct {
	core.QObject

	_ func(data int)                                                   `signal:"sendToQml"`
	_ func(pip int, vt int, rate int, peep int, fio2 int, mode string) `signal:"sendMonitor"`
	_ func(data string) string                                         `slot:"sendToGo"`
	_ func(ie int, pip int, bpm int, pmax int, peep int, fio2 int)     `slot:"sendPAC"`
}
