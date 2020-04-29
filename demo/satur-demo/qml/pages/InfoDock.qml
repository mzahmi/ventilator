import QtQuick 2.0
import QtQuick.Controls 2.0

Item {
    id: element
    width: 800
    height: 200
    property int progress: 50

    Rectangle {
        id: rectangle
        x: 0
        y: 0
        height: 50
        color: "#ffffff"
        radius: 5
        anchors.top: parent.top
        anchors.topMargin: 40
        anchors.right: parent.right
        anchors.rightMargin: 20
        anchors.left: parent.left
        anchors.leftMargin: 20

        Rectangle {

            id: rectangle_progress
            x: 0
            y: 0
            width: (progress*parent.width)/100 - 20
            // 50 is constant
            color: "#84a5c3"
            radius: 5
            anchors.leftMargin: 10
            anchors.bottomMargin: 10
            anchors.topMargin: 10
            anchors.bottom: parent.bottom
            anchors.left: parent.left
            anchors.top: parent.top
        }
    }




}

/*##^##
Designer {
    D{i:2;anchors_height:0;anchors_width:760;anchors_y:11;invisible:true}D{i:1;anchors_width:616;anchors_x:79;anchors_y:58}
}
##^##*/
