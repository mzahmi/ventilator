import QtQuick 2.0

Item {
    anchors.fill: parent
    Rectangle{
        anchors.fill: parent

        Text {
            id: element
            text: qsTr("Under Construction")
            anchors.fill: parent
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.pixelSize: 32
        }
    }

}
