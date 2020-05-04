package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

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

	go func() {
		for range time.NewTicker(time.Millisecond * 50).C {
			randnumber := rand.Intn(30)
			qmlBridge.SendToQml(randnumber)
		}
	}()

	app.Exec()
}

type QmlBridge struct {
	core.QObject

	_ func(data int)           `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`
}
