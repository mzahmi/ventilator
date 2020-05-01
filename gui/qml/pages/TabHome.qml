import QtQuick 2.0
import QtQuick 2.4
import QtQuick.Layouts 1.3

Item {
    RowLayout {
    anchors.top: parent.top
    anchors.right: parent.right
    anchors.bottom: parent.bottom
    anchors.left: parent.left
    anchors.topMargin: 0
    anchors.leftMargin: 177
    spacing: 30
    anchors.rightMargin: 177
    anchors.bottomMargin: 0
    transformOrigin: Item.Center

    GridLayout {
        rows: 1
        columns: 2
        Text {
            id: breathRate
            verticalAlignment: Text.AlignTop
            Connections {
                target: QmlBridge
            }
            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
            font.pixelSize: 48
        }

        Text {
            id: element1
            text: qsTr("BPM")
            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
            font.pixelSize: 20
            Layout.bottomMargin: 10
        }

        Text {
            id: element2
            color: "#636060"
            text: qsTr("BREATH RATE")
            font.pixelSize: 16
            font.bold: false
            Layout.columnSpan: 2
        }
    }

    GridLayout {
        rows: 1
        columns: 2
        Text {
            id: peep
            verticalAlignment: Text.AlignTop
            Connections {
                target: QmlBridge
            }
            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
            font.pixelSize: 48
        }

        Text {
            id: element4
            text: qsTr("cmH2O")
            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
            font.pixelSize: 20
            Layout.bottomMargin: 10
        }

        Text {
            id: element5
            color: "#636060"
            text: qsTr("PEEP")
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            font.pixelSize: 16
            font.bold: false
            Layout.columnSpan: 2
        }
    }

    GridLayout {
        rows: 1
        columns: 2
        RowLayout {
            id: rowLayout
            width: 100
            height: 100
            Text {
                id: minuteVolume
                text: "20"
                verticalAlignment: Text.AlignTop
                Connections {
                    target: QmlBridge
                }
                Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                font.pixelSize: 48
            }

            Text {
                id: element7
                text: qsTr("L")
                verticalAlignment: Text.AlignTop
                Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                font.pixelSize: 20
                Layout.bottomMargin: 10
            }
        }

        Text {
            id: element8
            color: "#636060"
            text: qsTr("MINUTE VOLUME")
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            font.pixelSize: 16
            font.bold: false
            Layout.columnSpan: 2
        }
    }
    }/*##^##
    Designer {
        D{i:0;anchors_y:160}D{i:8;anchors_x:174;anchors_y:80}D{i:6;anchors_width:800}
    }
    ##^##*/
}
