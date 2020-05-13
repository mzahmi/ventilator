import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "src/cards"
import "./material/qml/material"
import QtGraphicalEffects 1.0

Item {
    id: root
    height: 300
    width: moderow.width+moderow.spacing
    signal clicked()
    signal activated(string mode)
    signal stop()
    Component.onCompleted:{
        ModeSelect.modeSelected.connect(root.activated)
        ModeSelect.stopVent.connect(root.stop)
    }
    onActivated:{
        if (ModeSelect.mode==="Pressure A/C"){
            activateButton(preset1)
        }
    }
    onStop:{
        deactivateButton(preset1)
    }
    function activateButton(buttonID){
        buttonID.active = true
    }

    function deactivateButton(buttonID){
        buttonID.active = false
    }
    

    Flickable {
        id: flickable
        contentWidth: moderow.width+moderow.spacing
        anchors.fill: parent

        RowLayout {
            id: moderow
            y: 10
            anchors.leftMargin: 10
            anchors.left: parent.left
            spacing: 15

            Rectangle {
                id: preset1
                property string title:"Pressure A/C"
                property bool active: false
                property string breath:"-"
                property string trigger:"-"
                color: "#ffffff"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
                Card {
                    id: card1
                    anchors.fill: parent
                    raised: false
                    RaisedButton {
                        id: raisedbutton1
                        x: 0
                        y: 124
                        height: 32
                        color: preset1.active?"red":"#5677fc"
                        text: preset1.active?"Stop":"Start"
                        textColor: "#ffffff"
                        anchors.bottom: parent.bottom
                        anchors.left: parent.left
                        anchors.right: parent.right
                        rippleColor: "#deffffff"
                        onClicked: {
                            
                            if (raisedbutton1.text === "Stop"){
                                ModeSelect.stopVentilation()
                            } else{
                                root.clicked()
                            }
                        }
                    }

                    Text {
                        id: element3
                        x: 45
                        y: 8
                        text: qsTr("Pressure A/C")
                        font.pixelSize: 18
                    }

                    RowLayout {
                        y: 40
                        anchors.rightMargin: 20
                        anchors.leftMargin: 20
                        anchors.left: parent.left
                        anchors.right: parent.right
                        spacing: 5

                        Text {
                            id: b1
                            color: "#555555"
                            text: qsTr("Breath:")
                            font.pixelSize: 14
                        }

                        Text {
                            id: b1v
                            color: "#555555"
                            text: preset1.active?ModeSelect.breath:"-"
                            Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                            font.pixelSize: 14
                        }
                    }

                    RowLayout {
                        y: 62
                        anchors.rightMargin: 20
                        Text {
                            id: t1
                            color: "#555555"
                            text: qsTr("Trigger:")
                            font.pixelSize: 14
                        }

                        Text {
                            id: t1v
                            color: "#555555"
                            text: preset1.active?ModeSelect.trigger:"-"
                            Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                            font.pixelSize: 14
                        }
                        anchors.right: parent.right
                        spacing: 5
                        anchors.left: parent.left
                        anchors.leftMargin: 20
                    }

                }
            }

            Rectangle {
                id: preset2
                property bool active: false
                color: "#ffffff"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
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
                        text: "In development"
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

                        font.pixelSize: 18
                    }

                    RowLayout {
                        y: 40
                        anchors.rightMargin: 20
                        anchors.leftMargin: 20
                        anchors.left: parent.left
                        anchors.right: parent.right
                        spacing: 5

                        Text {
                            id: b2
                            color: "#555555"
                            text: qsTr("Breath:")
                            font.pixelSize: 14
                        }

                        Text {
                            id: b2v
                            color: "#555555"
                            text: preset2.active?ModeSelect.breath:"-"
                            Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                            font.pixelSize: 14
                        }
                    }

                    RowLayout {
                        y: 62
                        anchors.rightMargin: 20
                        Text {
                            id: t2
                            color: "#555555"
                            text: qsTr("Trigger:")
                            font.pixelSize: 14
                        }

                        Text {
                            id: t2v
                            color: "#555555"
                            text: preset2.active?ModeSelect.trigger:"-"
                            Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                            font.pixelSize: 14
                        }
                        anchors.right: parent.right
                        spacing: 5
                        anchors.left: parent.left
                        anchors.leftMargin: 20
                    }
                }
            }

            Rectangle {
                id: preset3
                color: "#ffffff"
                Layout.preferredHeight: 120
                Layout.preferredWidth: 194
                Card {
                    id: card2
                    anchors.fill: parent
                    raised: false
                    RaisedButton {
                        id: raisedbutton2
                        x: 0
                        y: 124
                        height: 32
                        color: "#b3b3b3"
                        text: "In development"
                        textColor: "#ffffff"
                        anchors.bottom: parent.bottom
                        anchors.left: parent.left
                        anchors.right: parent.right
                        rippleColor: "#deffffff"
                    }

                    Text {
                        id: title3
                        y: 7
                        text: qsTr("PSV")
                        anchors.left: parent.left
                        anchors.right: parent.right
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter

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

            Rectangle {
                id: preset4
                color: "#ffffff"
                Layout.preferredWidth: 194
                Card {
                    id: card3
                    anchors.fill: parent
                    RaisedButton {
                        id: raisedbutton3
                        x: 0
                        y: 124
                        height: 32
                        color: "#b3b3b3"
                        text: "In development"
                        anchors.left: parent.left
                        textColor: "#ffffff"
                        rippleColor: "#deffffff"
                        anchors.right: parent.right
                        anchors.bottom: parent.bottom
                    }

                    Text {
                        id: element7
                        y: 7
                        text: qsTr("P-SIMV")
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter
                        anchors.left: parent.left
                        anchors.right: parent.right

                        font.pixelSize: 18
                    }

                    Text {
                        id: element8
                        x: 63
                        y: 42
                        color: "#555555"
                        text: qsTr("50 / 32")
                        font.pixelSize: 18
                    }
                    raised: false
                }
                Layout.preferredHeight: 120
            }

            Rectangle {
                id: preset5
                color: "#ffffff"
                Layout.preferredWidth: 194
                Card {
                    id: card4
                    anchors.fill: parent
                    RaisedButton {
                        id: raisedbutton4
                        x: 0
                        y: 124
                        height: 32
                        color: "#b3b3b3"
                        text: "In development"
                        anchors.left: parent.left
                        textColor: "#ffffff"
                        rippleColor: "#deffffff"
                        anchors.right: parent.right
                        anchors.bottom: parent.bottom
                    }

                    Text {
                        id: element9
                        y: 7
                        text: qsTr("V-SIMV")
                        anchors.left: parent.left
                        anchors.right: parent.right
                        horizontalAlignment: Text.AlignHCenter

                        verticalAlignment: Text.AlignVCenter
                        font.pixelSize: 18
                    }

                    Text {
                        id: element10
                        x: 63
                        y: 42
                        color: "#555555"
                        text: qsTr("50 / 32")
                        font.pixelSize: 18
                    }
                    raised: false
                }
                Layout.preferredHeight: 120
            }
        }
    }
}

/*##^##
Designer {
    D{i:0;formeditorZoom:2}D{i:21;anchors_width:193;anchors_x:"-69"}D{i:20;anchors_x:70}
D{i:26;anchors_x:70}D{i:27;anchors_width:193;anchors_x:"-69"}D{i:32;anchors_x:70}
D{i:33;anchors_height:200;anchors_width:193;anchors_x:"-69";anchors_y:170}D{i:2;anchors_x:8}
}
##^##*/
