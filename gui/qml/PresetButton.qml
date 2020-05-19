import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "src/cards"
import "./material/qml/material"
import "./config.js"
as Config
import QtGraphicalEffects 1.0

Rectangle {
    id: root
    property string title: "Pressure A/C"
    property bool active: false
    property string breath: "-"
    property string trigger: "-"
    property bool development: true
    signal clicked()
    Card {
        id: card1
        anchors.fill: parent
        raised: false
        RaisedButton {
            id: raisedbutton1
            x: 0
            y: 124
            height: 32
            color: {
                if (!root.development) {
                    root.active ? "red" : "#5677fc"
                } else {
                    return Config.color_inactive
                }
            }

            text: {
                if (!root.development) {
                    root.active ? "Stop" : "Start"
                } else {
                    return "In Development"
                }
            }
            textColor: "#ffffff"
            anchors.bottom: parent.bottom
            anchors.left: parent.left
            anchors.right: parent.right
            rippleColor: "#deffffff"
            onClicked: {

                if (raisedbutton1.text === "Stop") {
                    ModeSelect.stopVentilation()
                } else if (!root.development) {
                    root.clicked()
                }
            }
        }

        Text {
            id: element3
            y: 8
            text: root.title
            anchors.left: parent.left
            anchors.right: parent.right
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.pixelSize: 18
        }

        RowLayout {
            y: 40
            anchors.rightMargin: 20
            anchors.leftMargin: 20
            anchors.left: parent.left
            anchors.right: parent.right
            spacing: 5

            Text {
                color: "#555555"
                text: qsTr("Breath:")
                font.pixelSize: 14
            }

            Text {
                color: "#555555"
                text: root.active ? ModeSelect.breath : "-"
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                font.pixelSize: 14
            }
        }

        RowLayout {
            y: 62
            anchors.rightMargin: 20
            Text {
                color: "#555555"
                text: qsTr("Trigger:")
                font.pixelSize: 14
            }

            Text {
                color: "#555555"
                text: root.active ? ModeSelect.trigger : "-"
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                font.pixelSize: 14
            }
            anchors.right: parent.right
            spacing: 5
            anchors.left: parent.left
            anchors.leftMargin: 20
        }

    }
}
