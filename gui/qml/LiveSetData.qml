import QtQuick 2.0
import "./src/variables/fontawesome.js"
as FontAwesome
import "./src/lists"
import "./config.js"
as Config
import QtQuick.Layouts 1.0
import "."
import QtQuick.Controls 2.1


Item {
    id: monitorrow
    property int pip: 25
    property int vt: 10
    property int rate: 15
    property int peep: 4
    property int fio2: 25
    property string mode: "PAC"
    width: 150
    height: 0

    RowLayout {
        x: -9
        y: 34
        spacing: 0
        anchors.right: parent.right
        anchors.leftMargin: 0
        ColumnLayout {
            id: column6
            Text {
                id: element17
                color: "#ffffff"
                text: monitorrow.fio2
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: Config.sidebar_livetext_size1
            }

            Text {
                id: element18
                color: "#ffffff"
                text: qsTr("FIO2")
                Layout.fillWidth: true
                styleColor: "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: Config.sidebar_livetext_size2
            }

            Text {
                id: element19
                color: "#f9f9f9"
                text: qsTr("%")
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: Config.sidebar_livetext_size3
            }
        }

        ColumnLayout {
            id: column7
            width: Config.sidebar_height / 2
            Text {
                id: element20
                color: "#ffffff"
                text: monitorrow.vt
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: Config.sidebar_livetext_size1
            }

            Text {
                id: element21
                color: "#ffffff"
                text: qsTr("Vt")
                Layout.fillWidth: true
                styleColor: "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: Config.sidebar_livetext_size2
            }

            Text {
                id: element22
                color: "#f9f9f9"
                text: qsTr("ml")
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: Config.sidebar_livetext_size3
            }
        }
        anchors.rightMargin: 0
        anchors.left: parent.left
    }

    RowLayout {
        x: 0
        y: 124
        anchors.rightMargin: 0
        anchors.leftMargin: 0
        spacing: 0
        anchors.left: parent.left
        anchors.right: parent.right
        ColumnLayout {
            id: column3
            Text {
                id: element8
                color: "#ffffff"
                text: monitorrow.rate
                font.pixelSize: Config.sidebar_livetext_size1
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }

            Text {
                id: element9
                color: "#ffffff"
                text: qsTr("Rate")
                font.pixelSize: Config.sidebar_livetext_size2
                styleColor: "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }

            Text {
                id: element10
                color: "#f9f9f9"
                text: qsTr("BPM")
                font.pixelSize: Config.sidebar_livetext_size3
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
        }

        ColumnLayout {
            id: column4
            width: Config.sidebar_height / 2
            Text {
                id: element11
                color: "#ffffff"
                text: monitorrow.peep
                font.pixelSize: Config.sidebar_livetext_size1
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }

            Text {
                id: element12
                color: "#ffffff"
                text: qsTr("PEEP")
                font.pixelSize: Config.sidebar_livetext_size2
                styleColor: "#ffffff"
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }

            Text {
                id: element13
                color: "#f9f9f9"
                text: qsTr("cmH2O")
                font.pixelSize: Config.sidebar_livetext_size3
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }
        }
    }




}

