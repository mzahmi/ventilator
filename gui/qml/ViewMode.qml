import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "./material/qml/material"
import "./config.js"
as Config
import "."

Item {
    id: name
    width: 650
    height: 480
    signal presetClicked(string mode)
    onPresetClicked: {
        if (mode === "Pressure A/C") {
            modePage.visible = false
            modePAC.visible = true
            modeVAC.visible = false
        }
        if (mode === "Volume A/C") {
            modePage.visible = false
            modeVAC.visible = true
            modePAC.visible = false
        }
    }
    signal stop()

    Component.onCompleted: {
        ModeSelect.stopVent.connect(name.stop)
    }

    onStop: {
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


        Item {
            id: modePage
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
                onClicked: {
                    console.log("not active")
                }
            }

            BaseLargeButton {
                id: volumeac3
                x: 12
                y: 38
                info: "Pressure control breaths at the set SIMV rate. Pressure/Flow trigger to assisted breath"
                title: "Pressure SIMV"
                onClicked: {
                    console.log("not active")
                }
            }

            BaseLargeButton {
                id: volumeac2
                x: -237
                y: 38
                active: true
                title: "Pressure A/C"
                info: "Suitable for passive, partially active and active patients with weak respiratory drive"
                onClicked: {
                    modePage.visible = false
                    modePAC.visible = true
                    modeVAC.visible = false
                }
            }

            BaseLargeButton {
                id: volumeac1
                x: 12
                y: -80
                title: "Volume SIMV"
                info: "Volume control breaths at the set SIMV rate. Pressure/Flow trigger to assisted breath"
                onClicked: {
                    console.log("not active")
                }
            }

            BaseLargeButton {
                id: volumeac
                x: -237
                y: -80
                active: true
                title: "Volume A/C"
                info: "Intended for patients who are passive or partially active"
                onClicked: {
                    modePage.visible = false
                    modePAC.visible = false
                    modeVAC.visible = true
                }
            }
        }

        ModePAC {
            id: modePAC
            visible: false
            anchors.left: parent.left
            anchors.right: parent.right

        }

        ModeVAC {
            id: modeVAC
            visible: false
            anchors.left: parent.left
            anchors.right: parent.right

        }

        Button {
            text: "back"
            onClicked: {
                modePAC.visible = false
                modeVAC.visible = false
                modePage.visible = true
            }

            visible: !modePage.visible

        }
    }
}



/*##^##
Designer {
    D{i:1;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
}
##^##*/
