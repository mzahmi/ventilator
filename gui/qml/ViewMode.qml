import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "./material/qml/material"
import "./config.js" as Config
import "./componentBCreation.js" as MyScript
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
    Component.onCompleted:{
        // MyScript.createButtonComponent(MyScript.toArray(ModeSelect.buttonList))
        ModeSelect.stopVent.connect(name.stop)
    }

    // when a stop signal appears
    // reset view
    onStop: {
        // hide button page
        flickableItems.visible=false
        // make page non interactive and move to top
        flickablePage.contentY = 0
        flickablePage.interactive=false
        // hide breath row
        rowBreath.visible=false
        // hide trigger
        rowTrigger.visible=false
        // show modes
        rowButtons.visible=true
        // remove all trigger buttons
        for(var i = rowTrigger.children.length; i > 0 ; i--) {
            console.log("destroying trigger")
            rowTrigger.children[i-1].height=0
        }
        // remove all breath buttons
        for(var j = rowBreath.children.length; j > 0 ; j--) {
            console.log("destroying breath")
            rowBreath.children[j-1].height=0
        }

        ModeSelect.status = "stop"
    }

    Flickable{
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

        Button{
            text: "back"
            onClicked: MyScript.backButton()
            
        }

//        Column {
//            id: rowButtons
//            y: 180
//            spacing: 10
//            anchors.rightMargin: 20
//            anchors.leftMargin: 20+this.spacing
//            anchors.right: parent.right
//            anchors.left: parent.left
//        }

        Item{
            id: rowButtons
            anchors.fill: parent

            BaseLargeButton{
                id: volumeac
                x: 76
                y: 166
                title:"Volume A/C"
                info: "Volume A/C"
            }

            BaseLargeButton {
                id: volumeac1
                x: 325
                y: 166
                title: "Volume SIMV"
                info: "Volume SIMV"
            }

            BaseLargeButton {
                id: volumeac2
                x: 76
                y: 284
                title: "Pressure A/C"
                info: "Pressure A/C"
            }

            BaseLargeButton {
                id: volumeac3
                x: 325
                y: 284
                info: "Pressure SIMV"
                title: "Pressure SIMV"
            }

            BaseLargeButton {
                id: volumeac4
                x: 76
                y: 364
                info: "Pressure Support (PSV)"
                title: "Pressure Support"
            }

            Text {
                id: element
                x: 76
                y: 255
                color: "#a8a8a8"
                text: qsTr("Pressure")
                font.bold: true
                font.pixelSize: 18
            }

            Text {
                id: element1
                x: 76
                y: 134
                color: "#a8a8a8"
                text: qsTr("Volume")
                font.pixelSize: 18
                font.bold: true
            }

        }

        Row {
            id: rowBreath
            y: 213
            spacing: 15
            anchors.rightMargin: 20
            anchors.leftMargin: 20+this.spacing
            anchors.right: parent.right
            anchors.left: parent.left
        }

        Row {
            id: rowTrigger
            y: 213
            spacing: 15
            anchors.rightMargin: 20
            anchors.leftMargin: 20+this.spacing
            anchors.right: parent.right
            anchors.left: parent.left
        }

        Item{
            visible:false
            id: flickableItems
        }



    }
}



/*##^##
Designer {
    D{i:5;anchors_height:200;anchors_width:200}D{i:6;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
D{i:8;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}D{i:9;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
D{i:10;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}D{i:11;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
D{i:1;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
}
##^##*/
