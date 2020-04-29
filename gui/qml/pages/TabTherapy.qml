import QtQuick 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import "MyConstants.js" as Constants


Item {
    id: element1
    width: 800
    height: 600
    property TextField chosenText

    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.top: parent.top
        anchors.bottom: parent.bottom
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.leftMargin: 0


    }

    Row {
        id: row
        anchors.fill: parent


        Flickable {
            id: flickable
            clip: true
            width: 500
            height: 650
            contentHeight: inputColumn.height+70
            contentWidth: 500
            anchors.left: parent.left
            anchors.leftMargin: 0
            anchors.bottom: parent.bottom
            anchors.top: parent.top
            anchors.topMargin: 0

            ColumnLayout {
                id: inputColumn
                x: 48
                y: 64
                spacing: 20

                RowLayout {
                    id: row1

                    ColumnLayout {
                        id: tidalColumn

                        Text {
                            id: tidalTitle
                            color: Constants.lightText
                            text: qsTr("Tidal Volume")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: tidalVolume
                            placeholderText: qsTr("Tidal Volume")
                            MouseArea{
                                onClicked: {
                                    chosenText = tidalVolume
                                    tidalVolume.focus= true
                                }
                                anchors.fill: parent
                            }
                        }

                    }

                    ColumnLayout {
                        id: rateColumn

                        Text {
                            id: rateTitle
                            color: Constants.lightText
                            text: qsTr("Rate")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: rate
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = rate
                                    rate.focus= true
                                }
                            }
                            placeholderText: qsTr("Flow Rate")
                        }
                    }
                }

                RowLayout {
                    id: row2
                    ColumnLayout {
                        id: tiColumn
                        Text {
                            id: tiTitle
                            color: Constants.lightText
                            text: qsTr("TI")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: tI
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = tI
                                    tI.focus= true
                                }
                            }
                            placeholderText: qsTr("TI")
                        }
                    }

                    ColumnLayout {
                        id: peakFlowColumn
                        Text {
                            id: peakFlowTitle
                            color: Constants.lightText
                            text: qsTr("Peak Flow")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: peakFlow
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = peakFlow
                                    peakFlow.focus= true
                                }
                            }
                            placeholderText: qsTr("Peak Flow")
                        }
                    }

                }

                RowLayout {
                    id: row3
                    ColumnLayout {
                        id: irColumn
                        Text {
                            id: irTitle
                            color: Constants.lightText
                            text: qsTr("IR")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: iR
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = iR
                                    iR.focus= true
                                }
                            }
                            placeholderText: qsTr("IR")
                        }
                    }

                    ColumnLayout {
                        id: eRColumn
                        Text {
                            id: eRTitle
                            color: Constants.lightText
                            text: qsTr("ER")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: eR
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = eR
                                    eR.focus= true
                                }
                            }
                            placeholderText: qsTr("ER")
                        }
                    }

                }

                RowLayout {
                    id: row4
                    ColumnLayout {
                        id: peepColumn
                        Text {
                            id: peepTitle
                            color: Constants.lightText
                            text: qsTr("PEEP")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: peep
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = peep
                                    peep.focus= true
                                }
                            }
                            placeholderText: qsTr("PEEP")
                        }
                    }

                    ColumnLayout {
                        id: fiO2Column
                        Text {
                            id: fiO2Title
                            color: Constants.lightText
                            text: qsTr("FIO2")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: fiO2
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = fiO2
                                    fiO2.focus= true
                                }
                            }
                            placeholderText: qsTr("FIO2")
                        }
                    }
                }

                RowLayout {
                    id: row5
                    width: 400
                    Layout.leftMargin: 0
                    Layout.fillWidth: true
                    ColumnLayout {
                        id: triggerTypeColumn
                        width: 200

                        Text {
                            id: triggerTypeTitle
                            color: Constants.lightText
                            text: qsTr("Trigger Type")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        ComboBox {
                            id: triggerType
                            width: 400
                            model: ["Time", "Pressure", "Flow"]
                            onActivated: {
                                console.log(currentIndex)
                            }
                        }


                    }

                    ColumnLayout {
                        id: trigSenseColumn
                        Layout.leftMargin: 60
                        Layout.fillWidth: false
                        Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                        Text {
                            id: trigSenseTitle
                            color: Constants.lightText
                            text: qsTr("Trigger Sensitivity")
                            font.family: "Tahoma"
                            font.pixelSize: 12
                        }

                        TextField {
                            id: trigSense
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    chosenText = trigSense
                                    trigSense.focus= true
                                }
                            }
                            placeholderText: qsTr("Trigger Sensitivity")
                        }
                    }

                }

                Row {
                    id: row6
                    width: 200
                    height: 400
                    spacing: 150
                    Layout.fillWidth: true

                    Button {
                        id: sendButton
                        text: qsTr("Send")
                        Connections{
                            target: QmlBridge
                        }

                        MouseArea {
                            anchors.leftMargin: 0
                            anchors.bottomMargin: 0
                            anchors.topMargin: 0
                            anchors.fill: parent
                            anchors.rightMargin: 0
                            onClicked: QmlBridge.sendToGo(tidalVolume.displayText, rate.displayText, tI.displayText, peakFlow.displayText, iR.displayText, eR.displayText, peep.displayText, fiO2.displayText, triggerType.currentIndex, trigSense.displayText)
                        }
                    }

                    Text {
                        id: info
                        font.bold: true
                        Connections{
                            target: QmlBridge
                            onSendInfo: info.text = data
                        }

                        font.pixelSize: 28
                    }
                }



            }


        }

        Rectangle {
            id: numpadBG
            color: Constants.mainbg
            anchors.right: parent.right
            anchors.left: flickable.right
            anchors.bottom: parent.bottom
            anchors.top: parent.top
            anchors.topMargin: 0

            GridLayout {
                columnSpacing: 0
                rowSpacing: 0
                anchors.bottom: parent.bottom
                anchors.top: parent.top
                anchors.topMargin: 50
                anchors.right: parent.right
                anchors.left: parent.left


                columns: 3
                rows: 2



                Button {
                    id: button7
                    width: 60
                    onClicked: chosenText.text = chosenText.text+7
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {

                        font.pointSize: 20
                        text: qsTr("7")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }

                Button {
                    id: button8
                    width: 60
                    onClicked: chosenText.text = chosenText.text+8
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {

                        font.pointSize: 20
                        text: qsTr("8")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }

                Button {
                    id: button9
                    width: 60
                    onClicked: chosenText.text = chosenText.text+9
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {

                        font.pointSize: 20
                        text: qsTr("9")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }


                Button {
                    id: button4
                    onClicked: chosenText.text = chosenText.text+4
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {

                        font.pointSize: 20
                        text: qsTr("4")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }



                Button {
                    id: button5
                    onClicked: chosenText.text = chosenText.text+5
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {

                        font.pointSize: 20
                        text: qsTr("5")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                    
                }


                Button {
                    id: button6
                    onClicked: chosenText.text = chosenText.text+6
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {

                        font.pointSize: 20
                        text: qsTr("6")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }

                Button {
                    id: button1
                    onClicked: chosenText.text = chosenText.text+1
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {
                        font.pointSize: 20
                        text: qsTr("1")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }

                Button {
                    id: button2
                    onClicked: chosenText.text = chosenText.text+2
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {
                        font.pointSize: 20
                        text: qsTr("2")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }

                Button {
                    id: button3
                    onClicked: chosenText.text = chosenText.text+3
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {
                        font.pointSize: 20
                        text: qsTr("3")
                        anchors.fill: parent
                        //opacity: enabled ? 1.0 : 0.3
                        //color: control.down ? "#17a81a" : "#21be2b"
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                }



                Button {
                    id: button0
                    onClicked: chosenText.text = chosenText.text+"0"
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {
                        text: qsTr("0")
                        verticalAlignment: Text.AlignVCenter
                        anchors.fill: parent
                        horizontalAlignment: Text.AlignHCenter
                        font.pointSize: 20
                        elide: Text.ElideRight
                    }
                }

                Button {
                    id: buttonDot
                    onClicked: chosenText.text = chosenText.text+"."
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {
                        text: qsTr(".")
                        verticalAlignment: Text.AlignVCenter
                        anchors.fill: parent
                        horizontalAlignment: Text.AlignHCenter
                        font.pointSize: 20
                        elide: Text.ElideRight
                    }
                }

                Button {
                    id: buttonDelete
                    onClicked: chosenText.text = chosenText.text.substring(0, chosenText.text.length - 1);
                    palette {
                        button: Constants.mainbg
                    }
                    contentItem: Text {
                        text: qsTr("<")
                        verticalAlignment: Text.AlignVCenter
                        anchors.fill: parent
                        horizontalAlignment: Text.AlignHCenter
                        font.pointSize: 20
                        elide: Text.ElideRight
                    }
                }
            }
        }
    }








    /*##^##
Designer {
    D{i:0;anchors_y:160}D{i:51;anchors_y:211}D{i:50;anchors_y:211}D{i:54;anchors_y:211}
D{i:53;anchors_y:211}D{i:49;anchors_y:211}D{i:56;anchors_y:211}D{i:55;anchors_y:211}
}
##^##*/
}
