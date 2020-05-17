import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "./material/qml/material"
import "./config.js"
as Config
import "./componentBCreation.js"
as MyScript
import "."

Item {
    id: name
    width: 650
    height: 480
    signal presetClicked()
    signal stop()
    onPresetClicked: {
        view.push(selectbreathe)
    }
    Component.onCompleted: {
        // MyScript.createButtonComponent(MyScript.toArray(ModeSelect.buttonList))
        ModeSelect.stopVent.connect(name.stop)
    }

    // when a stop signal appears
    // reset view
    onStop: {
        // hide button page
        flickableItems.visible = false
        // make page non interactive and move to top
        flickablePage.contentY = 0
        flickablePage.interactive = false
        // hide breath row
        rowBreath.visible = false
        // hide trigger
        rowTrigger.visible = false
        // show modes
        rowButtons.visible = true
        // remove all trigger buttons
        for (var i = rowTrigger.children.length; i > 0; i--) {
            rowTrigger.children[i - 1].height = 0
        }
        // remove all breath buttons
        for (var j = rowBreath.children.length; j > 0; j--) {
            rowBreath.children[j - 1].height = 0
        }

        ModeSelect.status = "stop"
    }

    Flickable {
        id: flickablePage
        interactive: false
        contentHeight: 500
        anchors.fill: parent
        Text {
            id: heading
            y: 60
            text: qsTr("Select Mode")
            horizontalAlignment: Text.AlignHCenter
            anchors.left: parent.left
            anchors.right: parent.right
            font.pointSize: 32
        }

        Button {
            text: "back"
            onClicked: MyScript.backButton()
            visible: !rowButtons.visible

        }

        Item {
            id: rowButtons
            anchors.verticalCenter: parent.verticalCenter
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.left: parent.horizontalCenter

            Text {
                id: element1
                x: -237
                y: -112
                color: "#a8a8a8"
                text: qsTr("Volume")
                font.pixelSize: 18
                font.bold: true
            }

            Text {
                id: element
                x: -237
                y: 9
                color: "#a8a8a8"
                text: qsTr("Pressure")
                font.bold: true
                font.pixelSize: 18
            }

            BaseLargeButton {
                id: volumeac4
                x: -237
                y: 118
                info: "Pressure Support ventilation mode is indicated for active patients only"
                title: "Pressure Support"
            }

            BaseLargeButton {
                id: volumeac3
                x: 12
                y: 38
                info: "Pressure control breaths at the set SIMV rate. Pressure/Flow trigger to assisted breath"
                title: "Pressure SIMV"
            }

            BaseLargeButton {
                id: volumeac2
                x: -237
                y: 38
                title: "Pressure A/C"
                info: "Suitable for passive, partially active and active patients with weak respiratory drive"
            }

            BaseLargeButton {
                id: volumeac1
                x: 12
                y: -80
                title: "Volume SIMV"
                info: "Volume control breaths at the set SIMV rate. Pressure/Flow trigger to assisted breath"
            }

            BaseLargeButton {
                id: volumeac
                x: -237
                y: -80
                title: "Volume A/C"
                info: "Intended for patients who are passive or partially active"
            }







        }

        Row {
            id: rowBreath
            y: 213
            spacing: 15
            anchors.rightMargin: 20
            anchors.leftMargin: 20 + this.spacing
            anchors.right: parent.right
            anchors.left: parent.left
        }

        Row {
            id: rowTrigger
            y: 213
            spacing: 15
            anchors.rightMargin: 20
            anchors.leftMargin: 20 + this.spacing
            anchors.right: parent.right
            anchors.left: parent.left
        }

        Item {
            visible: false
            id: flickableItems
        }



    }
}



/*##^##
Designer {
    D{i:1;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
}
##^##*/
