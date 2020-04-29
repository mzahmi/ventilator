import QtQuick 2.0

Item {
    id: breathText
    property int textSize: 32
    Text {
        id: element
        text: qsTr("Inhaling")
        font.bold: true
        anchors.fill: parent
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        font.pixelSize: textSize
    }

}
