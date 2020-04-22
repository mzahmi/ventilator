import QtQuick 2.0
import QtQuick.Window 2.12
import "pages"
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.13

Window {
    id: window
    visible: true
    width: 800
    height: 480
    title: qsTr("Hello World")

    color: "#ffffff"

    Rectangle {
        id: color_rectangle
        height: 245
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
            Layout.fillWidth: true
            progress: 28


        }

        MenuStack{
            Layout.fillHeight: true
            Layout.fillWidth: true

        }




    }
}

/*##^##
Designer {
    D{i:2;anchors_x:0;anchors_y:0}
}
##^##*/
