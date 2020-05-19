import QtQuick 2.0
import QtQuick.Window 2.1

Window {
    visible: true

    width: 1280
    height: 800
    visibility: fs ? "FullScreen" : "Windowed"
    title: qsTr("DFF Vent")
    Component.onCompleted: {
        console.log(Params.PEEP)
    }

    MainGo {

    }

}
