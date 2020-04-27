import QtQuick 2.0
import "pages"
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0

Item {
    id: window
    width: 800
    height: 480

    Rectangle {
        id: color_rectangle
        height: 230
        color: "#edf0f4"
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
    }



    ColumnLayout {
        anchors.bottom: parent.bottom
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
        spacing: 0

        StatusBar{
            Layout.fillWidth: true
            activeFocusOnTab: false


        }

        InfoDock{
            id: infoDock
            height: 120
            Layout.fillWidth: true
            Connections{
                target: QmlBridge
                onBreathRate: infoDock.progress = value
            }


        }

        Rectangle{
            height: 270
            Layout.fillHeight: true
            Layout.fillWidth: true

            Button {
                id: startSystem
                x: 202
                y: 140
                text: qsTr("Start")
                onClicked: QmlBridge.sendStart(bpm.text)
            }

            Button {
                id: stopSystem
                x: 432
                y: 140
                text: qsTr("Stop")
                onClicked: QmlBridge.sendStop("hello")
            }

            TextField {
                id: bpm
                x: 264
                y: 212
                placeholderText: "BPM"
            }


        }




    }

}


