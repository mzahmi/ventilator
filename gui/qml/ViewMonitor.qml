import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "src/cards"
import "./material/qml/material"


Item {
    id: root
    signal presetClicked()

    Rectangle {
        id: bg
        width: 800
        height: 480
        color: "#ffffff"
        anchors.fill: parent

        Rectangle {
            id: chartsarea
            height: 295
            color: "#ffffff"
            anchors.top: parent.top
            anchors.topMargin: 0
            anchors.left: parent.left
            anchors.leftMargin: 0
            anchors.right: parent.right
            anchors.rightMargin: 0

            BasicChart1{
                anchors.fill: parent

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

            ViewMonitorPreset{
                anchors.fill: parent
                onClicked: root.presetClicked()

            }
        }
    }
}



/*##^##
Designer {
    D{i:0;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46;autoSize:true;height:480;width:640}
D{i:2;anchors_height:300;anchors_width:300;anchors_x:162;anchors_y:29}D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/
