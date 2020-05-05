import QtQuick 2.0

Item {
    anchors.fill: parent
    Rectangle{
        anchors.fill: parent

        Rectangle {
            id: rectangle
            color: "#ffffff"
            anchors.fill: parent

            Text {
                id: element
                x: 86
                y: 63
                text: qsTr("Name: ")
                font.pixelSize: 12
            }

            Text {
                id: element1
                x: 86
                y: 90
                text: qsTr("Age: ")
                font.pixelSize: 12
            }

            Text {
                id: element2
                x: 86
                y: 123
                text: qsTr("Time on machine: ")
                font.pixelSize: 12
            }

            Text {
                id: element3
                x: 86
                y: 161
                text: qsTr("Preset values:")
                font.pixelSize: 12
            }
        }
    }

}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:2;anchors_height:200;anchors_width:200}
}
##^##*/
