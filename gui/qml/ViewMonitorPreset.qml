import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "src/cards"
import "./material/qml/material"
import QtGraphicalEffects 1.0

Item {
    id: root
    height: 300
    width: moderow.width + moderow.spacing
    signal clicked(string mode)
    signal activated(string mode)
    signal stop()
    // var presetList = [preset1, preset2, preset3, preset4, preset5]
    Component.onCompleted: {
        ModeSelect.modeSelected.connect(root.activated)
        ModeSelect.stopVent.connect(root.stop)
    }
    onActivated: {
        console.log("activated")
        if (ModeSelect.mode === "Pressure A/C") {
            preset1.active = true
        }
        if (ModeSelect.mode === "Volume A/C") {
            preset2.active = true
        }
    }
    onStop: {
        preset1.active = false
        preset2.active = false
    }


    Flickable {
        id: flickable
        contentWidth: moderow.width + moderow.spacing
        anchors.fill: parent

        RowLayout {
            id: moderow
            y: 10
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 20
            anchors.leftMargin: 10
            anchors.left: parent.left
            spacing: 15

            PresetButton {
                id: preset1
                title: "Pressure A/C"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
                color: "#ffffff"
                development: false
                onClicked: root.clicked(preset1.title)
            }

            PresetButton {
                id: preset2
                title: "Volume A/C"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
                color: "#ffffff"
                development: false
                onClicked: root.clicked(preset2.title)
            }

            PresetButton {
                id: preset3
                property bool active: false
                color: "#ffffff"
                title: "PSV"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
                onClicked: root.clicked()
            }
            PresetButton {
                id: preset4
                property bool active: false
                color: "#ffffff"
                title: "Volume SIMV"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
                onClicked: root.clicked()
            }

            PresetButton {
                id: preset5
                property bool active: false
                color: "#ffffff"
                title: "Pressure A/C"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
                onClicked: root.clicked()
            }

        }
    }
}

/*##^##
Designer {
    D{i:0;formeditorZoom:0.8999999761581421}D{i:6;anchors_x:45}D{i:16;anchors_x:45}D{i:21;anchors_width:193;anchors_x:"-69"}
D{i:20;anchors_x:70}D{i:26;anchors_x:45}D{i:27;anchors_width:193;anchors_x:"-69"}
D{i:32;anchors_x:70}D{i:33;anchors_height:200;anchors_width:193;anchors_x:"-69";anchors_y:170}
D{i:2;anchors_x:8}
}
##^##*/
