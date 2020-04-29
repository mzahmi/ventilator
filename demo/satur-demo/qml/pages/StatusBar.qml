import QtQuick 2.4
import QtQuick.Controls 2.0
import "icons"
import QtQuick.Layouts 1.0

StatusBarForm {
    id: statusBar
    height: 40

    anchors.top: parent.top
    anchors.topMargin: 0


    RowLayout {
        x: 680
        y: 4
        anchors.verticalCenterOffset: 1
        anchors.verticalCenter: parent.verticalCenter
        anchors.right: parent.right
        anchors.rightMargin: 8
        spacing: 8



        Text {
            id: element
            text: qsTr("SMIV-VOLT")
            font.pixelSize: 18
        }

        FontAwesomeIcon{
            width: 32
            height: 32
            iconSize: 24
            iconStyle: "\uf071"

        }

        FontAwesomeIcon{
            width: 32
            height: 32
            iconSize: 24
            iconStyle: "\uf076"

        }


        BatteryIcon{
            width: 32
            height: 32
            Layout.topMargin: 3
            Layout.preferredHeight: 32
            Layout.preferredWidth: 32
            iconSize: 24
            level: 1
        }



    }

}
