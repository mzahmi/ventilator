import QtQuick 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.1

Item {
    id: element1
    width: 800
    height: 300



    Rectangle {
        property TextEdit chosenText
        id: rectangle
        color: "#ffffff"
        anchors.top: parent.top
        anchors.topMargin: 50
        anchors.bottom: parent.bottom
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.leftMargin: 0
        Flickable{
            contentHeight: 1200
            contentWidth: 800
            anchors.fill: parent

        Button {
            id: button
            x: 350
            y: 192
            text: qsTr("Send")
            MouseArea
            {
                anchors.fill: parent
                onClicked: QmlBridge.sendToGo(tidalVolume.text, flowRate.text, tI.text, iR.text, eR.text, peakFlow.text, peep.text, fiO2.text, trigSense.text)
            }
        }


        GridLayout {
            x: 95
            y: 55
            rows: 3
            columns: 3

            TextField {
                id: tidalVolume
                placeholderText: qsTr("Tidal Volume")
                MouseArea{
                    onClicked: {
                        console.log("hi")
                        tidalVolume.focus = true
                        chosenText = tidalVolume}
                    anchors.fill: parent

                }
            }

            TextField {
                id: flowRate
                placeholderText: qsTr("Flow Rate")
            }

            TextField {
                id: tI
                placeholderText: qsTr("TI")
            }

            TextField {
                id: iR
                placeholderText: qsTr("IR")
            }

            TextField {
                id: eR
                placeholderText: qsTr("ER")
            }

            TextField {
                id: peakFlow
                placeholderText: qsTr("Peak Flow")
            }

            TextField {
                id: peep
                placeholderText: qsTr("PEEP")
            }

            TextField {
                id: fiO2
                placeholderText: qsTr("FIO2")
            }

            TextField {
                id: trigSense
                placeholderText: qsTr("Trigger Sensitivity")
            }
        }

        Text {
            id: element
            x: 95
            y: 27
            text: qsTr("Volume AC")
            font.pixelSize: 18
        }

        Button {
            id: button1
            x: 112
            y: 192
            text: qsTr("1")
            onClicked: {
                chosenText.text = 1
            }
        }


    }

}






    /*##^##
Designer {
    D{i:0;anchors_y:160}
}
##^##*/
}
