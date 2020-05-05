import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.0
import "src/cards"
import "./material/qml/material"

Item {

    Rectangle {
        id: bg
        anchors.fill: parent

    Rectangle {
        id: rectangle
        x: 223
        y: 10
        width: 194
        height: 120
        color: "#ffffff"
        Card {
            id: card
            anchors.fill: parent
            raised: false
            RaisedButton {
                id: raisedbutton
                x: 0
                y: 124
                height: 32
                color: "#b3b3b3"
                text: "Active"
                textColor: "#ffffff"
                anchors.bottom: parent.bottom
                anchors.left: parent.left
                anchors.right: parent.right
                rippleColor: "#deffffff"
            }

            Text {
                id: element1
                x: 45
                y: 8
                text: qsTr("VOLUME A/C")
                ToolSeparator {
                    id: toolSeparator
                    x: -45
                    y: 19
                    width: 193
                    height: 13
                    orientation: Qt.Horizontal
                }
                font.pixelSize: 18
            }

            Text {
                id: element2
                x: 63
                y: 42
                color: "#555555"
                text: qsTr("750 / 20")
                font.pixelSize: 18
            }
        }
    }

    Rectangle {
        id: rectangle2
        x: 8
        y: 10
        width: 194
        height: 120
        color: "#ffffff"
        Card {
            id: card1
            anchors.fill: parent
            raised: false
            RaisedButton {
                id: raisedbutton1
                x: 0
                y: 124
                height: 32
                color: "#5677fc"
                text: "Start"
                textColor: "#ffffff"
                anchors.bottom: parent.bottom
                anchors.left: parent.left
                anchors.right: parent.right
                rippleColor: "#deffffff"
            }

            Text {
                id: element3
                x: 45
                y: 8
                text: qsTr("Pressure A/C")
                ToolSeparator {
                    id: toolSeparator1
                    x: -45
                    y: 19
                    width: 193
                    height: 13
                    orientation: Qt.Horizontal
                }
                font.pixelSize: 18
            }

            Text {
                id: element4
                x: 63
                y: 42
                color: "#555555"
                text: qsTr("+20 / -20")
                font.pixelSize: 18
            }
        }

        MouseArea {
            anchors.fill: parent
            Connections {
                target: QmlBridge
            }
        }
    }

    Rectangle {
        id: rectangle3
        x: 438
        y: 10
        width: 194
        height: 120
        color: "#ffffff"
        Card {
            id: card2
            anchors.fill: parent
            raised: false
            RaisedButton {
                id: raisedbutton2
                x: 0
                y: 124
                height: 32
                color: "#5677fc"
                text: "Start"
                textColor: "#ffffff"
                anchors.bottom: parent.bottom
                anchors.left: parent.left
                anchors.right: parent.right
                rippleColor: "#deffffff"
            }

            Text {
                id: element5
                x: 70
                y: 7
                text: qsTr("SIMV")
                ToolSeparator {
                    id: toolSeparator2
                    x: -69
                    y: 20
                    width: 193
                    height: 13
                    orientation: Qt.Horizontal
                }
                font.pixelSize: 18
            }

            Text {
                id: element6
                x: 63
                y: 42
                color: "#555555"
                text: qsTr("50 / 32")
                font.pixelSize: 18
            }
        }
    }

    PageIndicator {
        id: pageIndicator
        x: 296
        y: 140
        count: 3
    }
    }/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}
}
##^##*/

}
