import QtQuick 2.0
import "pages"
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0

Item {
    id: window
    width: 800
    height: 460

    Rectangle {
        id: color_rectangle
        height: 230
        color: "#edf0f4"
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
    }


    ColumnLayout {
        anchors.bottom: parent.bottom
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
        spacing: 0




        StatusBar{
            Layout.fillWidth: true
            activeFocusOnTab: false


        }

        InfoDock{
            height: 100
            Layout.fillWidth: true


        }
        MenuStack{
            height: 270
            Layout.fillHeight: true
            Layout.fillWidth: true

        }
    }

}

/*##^##
Designer {
    D{i:1;anchors_height:245}D{i:2;anchors_x:0;anchors_y:0}
}
##^##*/
