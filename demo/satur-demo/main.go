/*
singal : from go to qml
slot : from qml to go
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"periph.io/x/periph/conn/gpio"
)

var (
	slider time.Duration = 0
)

// QmlBridge ... Custom bridge
type QmlBridge struct {
	core.QObject

	_ func(data string) `slot:"sendStart"`
	_ func(data string) `slot:"sendStop"`
	_ func(value int)   `signal:"breathRate"`

	// cretes onSendTime in qml
	_ func(data string) `signal:"sendTime"`
}

// creates variables that will be linked the bridge struct
var (
	// qmlBridge    = NewQmlBridge(nil)
	qmlBridge = NewQmlBridge(nil)
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
	// location of qml file
	view.SetSource(core.NewQUrl3("qrc:/qml/MainGo.qml", 0))

	// time counter
	// counter := 0

	// START BUTTON
	qmlBridge.ConnectSendStart(func(data string) {
		fmt.Println(fmt.Println("Start ", data))
		StartClicked()
		f, _ := strconv.ParseFloat(data, 8)
		CalculateBCT(float32(f))
	})

	// STOP BUTTON
	qmlBridge.ConnectSendStop(func(data string) {
		fmt.Println(fmt.Println("Stop", data))
		StopClicked()
	})

	// initiate host
	InitiateHost()

	go func() {
		for {
			qmlBridge.BreathRate(int(slider))
			//fmt.Println(readFromFile2())
			//time.Sleep(time.Second)
			//counter++
		}
	}()

	// control loop
	go func() {
		for {

			time.Sleep(time.Millisecond * 200)
			for !StartBool {
				slider = 0
				fmt.Println("Entered endless loop")
				highBool := true

				for start := time.Now(); time.Since(start) < (time.Millisecond * time.Duration(TI*1000)); {
					time.Sleep(time.Millisecond * 50)
					slider = 100 * time.Since(start) / (time.Millisecond * time.Duration(TI*1000))
					fmt.Println(slider)
					if highBool {
						MIns.Out(gpio.High)
						highBool = false
					}
				}

				MIns.Out(gpio.Low)

				for start := time.Now(); time.Since(start) < (time.Millisecond * time.Duration(TE*1000)); {
					time.Sleep(time.Millisecond * 50)
					slider = 100 - (100 * time.Since(start) / (time.Millisecond * time.Duration(TE*1000)))
					fmt.Println(slider)
					if !highBool {
						MExp.Out(gpio.High)
						highBool = true
					}
				}

				MExp.Out(gpio.Low)
			}
		}
	}()

	// show the view
	view.Show()

	// keep program running until exit window
	gui.QGuiApplication_Exec()
}

func readFromFile2() int {
	returnString := "20"

	data, err := ioutil.ReadFile("test2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		returnString = "No file found"
	} else {
		returnString = string(data)
	}
	i1, _ := strconv.Atoi(returnString)

	return i1
}
