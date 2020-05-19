import QtQuick 2.0
import QtQuick.Window 2.1

Window {
    visible: true

    width: 800
    height: 480
    visibility: fs ? "FullScreen" : "Windowed"
    title: qsTr("DFF Vent")

    MainGo {

    }

}
