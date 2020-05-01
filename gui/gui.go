package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

// QmlBridge ... Custom bridge
type QmlBridge struct {
	core.QObject

	_ func(data string)                                                                                                                                             `signal:"sendToQml"`
	_ func(tidalVolume float32, rate float32, ti float32, peakFlow float32, ir float32, er float32, peep float32, fio2 float32, triggerType int, trigSense float32) `slot:"sendToGo"`
	_ func(data string)                                                                                                                                             `signal:"sendInfo"`
	_ func(data string)                                                                                                                                             `signal:"sendMinutevolume"`
	_ func(mode string)                                                                                                                                             `slot:"sendModeSelected"`
}

// InitView Initiates quickview and returns a view
func InitView() *quick.QQuickView {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)

	// creates a quickview type and
	// connects the bridge struct with thebridge variable
	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	// view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)

	// location of qml file
	view.SetSource(core.NewQUrl3("qrc:/qml/MainGo.qml", 0))
	return view
}

// creates variables that will be linked the bridge struct
var (
	// qmlBridge    = NewQmlBridge(nil)
	qmlBridge = NewQmlBridge(nil)
)
