import QtQuick 2.0
import QtQuick.Layouts 1.0
import QtQuick.Controls 2.1
import "./config.js"
as Config

Item {
    anchors.fill: parent
    Rectangle {
        id: rectangle
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 0
        anchors.fill: parent

        Text {
            id: element
            y: 41
            color: "#646464"
            text: qsTr("Patient Information")
            anchors.leftMargin: 40
            anchors.left: parent.left
            anchors.right: parent.right
            font.pixelSize: 28
        }

        Text {
            id: element1
            x: 40
            y: 86
            color: "#6e6e6e"
            text: qsTr("Step 1:")
            font.pixelSize: 16
        }

        Text {
            id: element2
            x: 97
            y: 86
            text: qsTr("Tell us about yourself")
            font.bold: true
            font.pixelSize: 16
        }

        Text {
            id: element3
            x: 40
            y: 157
            color: "#6c6c6c"
            text: qsTr("What is your gender?")
            font.pixelSize: 15
        }

        Rectangle {
            id: male
            property bool active: false
            x: 40
            y: 191
            width: 106
            height: 41
            color: "#ffffff"
            radius: 8
            border.color: active ? Config.color_primary : "grey"


            Text {
                y: 15
                color: male.active ? Config.color_primary : "grey"
                text: qsTr("Male")
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                anchors.right: parent.right
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }

            MouseArea {
                id: mouseAreaMale
                anchors.fill: parent
                onClicked: {
                    male.active = true
                    female.active = false
                }
            }
        }

        Rectangle {
            id: female
            property bool active: false
            x: 173
            y: 191
            width: 106
            height: 41
            color: "#ffffff"
            radius: 8
            border.color: active ? Config.color_primary : "grey"

            Text {
                y: 15
                color: female.active ? Config.color_primary : "grey"
                text: qsTr("Female")
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                anchors.right: parent.right
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }

            MouseArea {
                id: mouseAreaFemale
                anchors.fill: parent
                onClicked: {
                    female.active = true
                    male.active = false
                }
            }
        }

        Text {
            id: element4
            x: 40
            y: 277
            color: "#6c6c6c"
            text: qsTr("How old are you?")
            font.pixelSize: 15
        }

        SpinBox {
            id: spinBox
            x: 40
            y: 314
        }


    }

}

/*##^##
Designer {
    D{i:0;autoSize:true;formeditorZoom:1.75;height:480;width:640}D{i:2;anchors_x:32}D{i:7;anchors_height:10;anchors_x:8;anchors_y:16}
D{i:10;anchors_x:45}D{i:11;anchors_height:100;anchors_width:100}
}
##^##*/
