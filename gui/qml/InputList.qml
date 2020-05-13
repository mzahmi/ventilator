import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "./config.js" as Config
import "./componentCreation.js" as SlrScript
import "./componentBCreation.js" as BtnScript

Item {
    id: name
    height: 800
    property var inputList
    property int contentHeight: flickable.contentHeight
    property int ieratio: 1
    Component.onCompleted: {
        var componentNumber = SlrScript.getComponents(inputList)
    }

    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.fill: parent

        Flickable {
            id: flickable
            anchors.rightMargin: 0
            anchors.leftMargin: 0
            contentHeight: flickableItems.children.length * 120 + 100
            anchors.fill: parent
            Component.onCompleted:{
                BtnScript.inputListHeight=contentHeight
            }

            Column {
                id: column
                anchors.topMargin: 40
                anchors.fill: parent

                // Text {
                //     id: grouptitle
                //     x: 80
                //     y: 15
                //     text: qsTr("Pressure Assist/Control")
                //     font.bold: false
                //     verticalAlignment: Text.AlignTop
                //     horizontalAlignment: Text.AlignHCenter
                //     font.pixelSize: 28
                // }

                Item {
                    id: flickableItems
                }
            }

            Rectangle {
                id: buttonSubmit
                x: 267
                y: flickableItems.children[flickableItems.children.length - 1].y + 125
                width: 110
                height: 38
                color: Config.color_dark
                MouseArea {
                    anchors.fill: parent
                    onReleased: {
                        SlrScript.getComponentsValues(flickableItems)
                        ModeSelect.startVentilation()
                        //foo.test_slot(name.ieratio.toString()+ " "+slider2.value.toString()+ " "+slider3.value.toString()+ " "+slider4.value.toString()+ " "+slider7.value.toString()+ " "+slider8.value.toString())
                        //foo.test_slot(flickableItems.children[i].name, flickableItems.children[i].value)
                    }
                }

                Text {
                    id: element
                    text: qsTr("SUBMIT")
                    font.bold: true
                    color: "white"
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                    anchors.fill: parent
                    font.pixelSize: 12
                }
            }
        }
    }
}

/*##^##
Designer {
    D{i:4;anchors_x:38;anchors_y:24}D{i:8;anchors_height:8;anchors_width:8;anchors_x:6;anchors_y:6}
D{i:7;anchors_height:8;anchors_width:8;anchors_x:6;anchors_y:6}D{i:9;anchors_height:8;anchors_width:8;anchors_x:6;anchors_y:6}
D{i:2;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/

