import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "src/cards"
import "./material/qml/material"


Item {
    id: root
    signal presetClicked()
    anchors.fill: parent



    Rectangle {
        id: chartsarea
        height: parent.height * 2 / 3
        color: "#ffffff"
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0

        BasicChart1 {
            id: chart1
            property bool active: true
            visible: active ? true : false
            anchors.fill: parent
            anchors.topMargin: 10

        }

        BasicChart2 {
            id: chart2
            property bool active: false
            anchors.fill: parent
            anchors.topMargin: 10
            visible: active ? true : false
        }

        BasicChart3 {
            id: chart3
            property bool active: false
            anchors.fill: parent
            anchors.topMargin: 10
            visible: active ? true : false
        }

        RowLayout {
            x: 188
            y: 277
            spacing: 15

            MouseArea {
                id: pressurema
                Layout.preferredHeight: 35
                Layout.preferredWidth: 78
                onClicked: {
                    chart1.active = true
                    chart2.active = false
                    chart3.active = false
                }

                Text {
                    text: "Pressure"
                    font.bold: chart1.active ? true : false
                    color: chart1.active ? "blue" : "black"
                    anchors.fill: parent
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                }
            }

            MouseArea {
                id: volumema
                Layout.preferredHeight: 35
                Layout.preferredWidth: 78
                onClicked: {
                    chart2.active = true
                    chart1.active = false
                    chart3.active = false
                }

                Text {
                    text: "Volume"
                    verticalAlignment: Text.AlignVCenter
                    font.bold: chart2.active ? true : false
                    color: chart2.active ? "blue" : "black"
                    horizontalAlignment: Text.AlignHCenter
                    anchors.fill: parent
                }
            }

            MouseArea {
                id: flowma
                Layout.preferredHeight: 35
                Layout.preferredWidth: 78
                onClicked: {
                    chart2.active = false
                    chart1.active = false
                    chart3.active = true
                }

                Text {
                    text: "Flow"
                    font.bold: chart3.active ? true : false
                    color: chart3.active ? "blue" : "black"
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                    anchors.fill: parent
                }
            }
        }

    }

    Rectangle {
        id: rectangle1
        color: "#ffffff"
        anchors.top: chartsarea.bottom
        anchors.bottom: parent.bottom
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.leftMargin: 0

        ViewMonitorPreset {
            anchors.fill: parent
            onClicked: root.presetClicked()

        }
    }

}



/*##^##
Designer {
    D{i:0;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46;autoSize:true;height:480;width:640}
D{i:2;anchors_height:295;anchors_width:300;anchors_x:162;anchors_y:29}D{i:3;anchors_height:500}
D{i:9;anchors_height:500}D{i:1;anchors_height:200;anchors_width:200}D{i:12;anchors_height:500}
}
##^##*/
