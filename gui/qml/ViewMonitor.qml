import QtQuick 2.0
import QtQuick.Controls 2.13
import QtCharts 2.3
import QtQuick.Layouts 1.0
import "src/cards"
import "./material/qml/material"

Item {
    Rectangle {
        id: bg
        color: "#ffffff"
        anchors.fill: parent

        Rectangle {
            id: chartsarea
            height: 200
            color: "#ffffff"
            anchors.top: parent.top
            anchors.topMargin: 0
            anchors.left: parent.left
            anchors.leftMargin: 0
            anchors.right: parent.right
            anchors.rightMargin: 0

            Chart1{
                anchors.fill: parent

            }
        }

        Rectangle {
            id: rectangle1
            color: "#ffffff"
            anchors.bottom: parent.bottom
            anchors.top: chartsarea.bottom
            anchors.right: parent.right
            anchors.left: parent.left
            anchors.leftMargin: 0

            Rectangle {
                id: rectangle
                x: 223
                y: 99
                width: 194
                height: 120
                color: "#ffffff"

                Card{
                    id: card
                    raised: false
                    anchors.fill: parent

                    RaisedButton {
                        id: raisedbutton
                        x: 0
                        y: 124
                        height: 32
                        text: "Active"
                        anchors.bottom: parent.bottom
                        anchors.right: parent.right
                        anchors.left: parent.left
                        color: "#b3b3b3"
                        textColor: "white"
                        rippleColor: "#deffffff"
                    }

                    Text {
                        id: element1
                        x: 45
                        y: 8
                        text: qsTr("VOLUME A/C")
                        font.pixelSize: 18

                        ToolSeparator {
                            id: toolSeparator
                            x: -45
                            y: 19
                            width: 193
                            height: 13
                            orientation: Qt.Horizontal
                        }
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

            Text {
                id: element
                x: 118
                y: 0
                width: 421
                height: 59
                color: "#484848"
                text: qsTr("Select a mode to run the ventilator on, preset values can be customized")
                wrapMode: Text.WordWrap
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 20
            }

            Rectangle {
                id: rectangle2
                x: 8
                y: 99
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
                        anchors.bottom: parent.bottom
                        anchors.right: parent.right
                        rippleColor: "#deffffff"
                        anchors.left: parent.left
                        textColor: "#ffffff"
                    }

                    Text {
                        id: element3
                        x: 45
                        y: 8
                        text: qsTr("Pressure A/C")
                        font.pixelSize: 18
                        ToolSeparator {
                            id: toolSeparator1
                            x: -45
                            y: 19
                            width: 193
                            height: 13
                            orientation: Qt.Horizontal
                        }
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
            }

            Rectangle {
                id: rectangle3
                x: 438
                y: 99
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
                        anchors.right: parent.right
                        anchors.bottom: parent.bottom
                        rippleColor: "#deffffff"
                        anchors.left: parent.left
                        textColor: "#ffffff"
                    }

                    Text {
                        id: element5
                        x: 70
                        y: 7
                        text: qsTr("SIMV")
                        font.pixelSize: 18
                        ToolSeparator {
                            id: toolSeparator2
                            x: -69
                            y: 20
                            width: 193
                            height: 13
                            orientation: Qt.Horizontal
                        }
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
                y: 235
                count: 3
            }
        }
    }
}



/*##^##
Designer {
    D{i:0;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46;autoSize:true;height:480;width:640}
D{i:2;anchors_height:300;anchors_width:300;anchors_x:162;anchors_y:29}D{i:8;anchors_width:200}
D{i:10;anchors_height:400;anchors_width:153;anchors_x:87;anchors_y:0}D{i:11;anchors_height:156;anchors_width:198;anchors_x:0;anchors_y:0}
D{i:17;anchors_height:400;anchors_width:153;anchors_x:87;anchors_y:0}D{i:23;anchors_height:400;anchors_width:153;anchors_x:87;anchors_y:0}
D{i:18;anchors_height:156;anchors_width:198;anchors_x:0;anchors_y:0}D{i:24;anchors_height:156;anchors_width:198;anchors_x:0;anchors_y:0}
D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/
