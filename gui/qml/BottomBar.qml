import QtQuick 2.0
import QtQuick.Controls 2.1

Item {
    id: element
    width: 650
    height: 100

    Row {
        id: row
        anchors.fill: parent

        BasicSetButton {
            id: basicSetButton
            anchors.top: parent.top
            anchors.bottom: parent.bottom
            width: parent.width/4
            anchors.bottomMargin: 0
            anchors.topMargin: 0
        }

        BasicSetButton {
            id: basicSetButton1
            width: parent.width/4
            anchors.topMargin: 0
            anchors.top: parent.top
            anchors.bottomMargin: 0
            anchors.bottom: parent.bottom
        }
    }



}
