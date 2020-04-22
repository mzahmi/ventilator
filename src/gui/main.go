/*
singal : from go to qml
slot : from qml to go
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

// // QmlBridge ... test bridge
// type QmlBridge struct {
// 	core.QObject

// 	// send to go
// 	_ func(data string)        `signal:"sendToQml"`
// 	_ func(data string) string `slot:"sendToGo"`
// }

// QmlBridge ... Custom bridge
type QmlBridge struct {
	core.QObject

	_ func(data string)        `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`

	// cretes onSendTime in qml
	_ func(data string) `signal:"sendTime"`
}

// ButtonBridge ... Custom bridge
type ButtonBridge struct {
	core.QObject

	_ func(data int)           `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`
}

// creates variables that will be linked the bridge struct
var (
	// qmlBridge    = NewQmlBridge(nil)
	qmlBridge    = NewQmlBridge(nil)
	buttonBridge = NewQmlBridge(nil)
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	// creates a quickview type and
	// connects the bridge struct with thebridge variable
	var view = quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	// view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.RootContext().SetContextProperty("ButtonBridge", buttonBridge)
	// location of qml file
	view.SetSource(core.NewQUrl3("qrc:/qml/mainGo.qml", 0))

	// time counter
	counter := 0

	// button function when clicked
	buttonBridge.ConnectSendToGo(func(data string) string {
		f, err := os.OpenFile("log.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(strconv.Itoa(counter) + "\n"); err != nil {
			log.Println(err)
		}
		return "hello from go"
	})

	// show time
	go func() {
		for t := range time.NewTicker(time.Second * 1).C {
			qmlBridge.SendTime(t.Format(time.ANSIC))
		}
	}()

	// read from a file every second
	go func() {
		i := 0
		for i <= 100 {
			qmlBridge.SendToQml(readFromFile())
			time.Sleep(time.Second)
			counter++
		}
	}()

	// show the view
	view.Show()

	// keep program running until exit window
	gui.QGuiApplication_Exec()
}

func readFromFile() string {
	returnString := ""

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		returnString = "No file found"
	} else {
		returnString = string(data)
	}
	return returnString
}
