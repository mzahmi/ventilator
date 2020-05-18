import QtQuick 2.0
import QtQuick.Controls 2.1

Item {
    id: root
    property string name: "PEEP"
    property int value: 10
    property string unit: "cmH2O"
    property int min: 0
    property int min: 50


    anchors.fill: parent

    Text {
        id: value
        y: 37
        text: root.value
        anchors.verticalCenter: parent.verticalCenter
        anchors.right: parent.right
        anchors.left: parent.left
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
        font.pixelSize: 18
    }

    Text {
        id: min
        y: 42
        text: root.min
        anchors.verticalCenter: parent.verticalCenter
        anchors.left: parent.left
        anchors.leftMargin: 8
        font.pixelSize: 12
    }

    Text {
        id: max
        y: 42
        height: 17
        text: root.max
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.rightMargin: 8
        anchors.verticalCenter: parent.verticalCenter
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignRight
        font.pixelSize: 12
    }

    Text {
        id: name
        x: 88
        text: root.name
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.top: parent.top
        anchors.topMargin: 8
        font.pixelSize: 14
    }

    Text {
        id: units
        x: 88
        y: 8
        text: root.units
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 8
        anchors.horizontalCenterOffset: 0
        font.pixelSize: 12
        anchors.horizontalCenter: parent.horizontalCenter
    }
}
