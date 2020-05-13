import QtQuick 2.0
import QtQuick.Controls 2.0
import "./componentBCreation.js" as BtnScript

Item {
    id: root
    width: rectangle.width
    height: rectangle.height
    property string title:"Volume A/C"
    property string info:"In Development"

    
    Rectangle {
        id: rectangle
        x: 0
        y: 0
        width: 233
        height: 74
        color: "#ffffff"
        radius: 10
        border.color: "#8e8e8e"





        Text {
            color: "#000000"
            text: root.title
            font.bold: true
            anchors.rightMargin: 0
            anchors.top: parent.top
            anchors.topMargin: 10
            horizontalAlignment: Text.AlignLeft
            anchors.left: parent.left
            anchors.leftMargin: 20
            anchors.right: parent.right
            font.pixelSize: 16
        }

        MouseArea{
            id: mousearea
            width: 70
            anchors.right: parent.right
            anchors.bottom: parent.bottom
            anchors.top: parent.top
            onClicked: BtnScript.createComponent(root.title)
        }

        Rectangle {
            id: rectangle1
            x: 174
            y: 27
            width: 49
            height: 20
            color: "#3247ef"
            radius: 5
            anchors.right: parent.right
            anchors.rightMargin: 10

            Text {
            id: element2
            color: "#ffffff"
            text: qsTr("Select")
            anchors.leftMargin: 0
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            anchors.fill: parent
            font.pixelSize: 12
            }

        }

        Text {
            id: element1
            height: 17
            color: "#535353"
            text: root.info
            anchors.right: rectangle1.left
            anchors.leftMargin: 20
            anchors.topMargin: 34
            anchors.top: parent.top
            font.pixelSize: 12
            horizontalAlignment: Text.AlignLeft
            anchors.left: parent.left
        }

    }

}

/*##^##
Designer {
    D{i:0;formeditorZoom:1.75}
}
##^##*/
