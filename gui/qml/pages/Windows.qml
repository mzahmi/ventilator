import QtQuick 2.12
import QtQuick.Controls 2.5
import QtQuick.Layouts 1.14

Item {
    StackLayout {
        anchors.top: parent.top
        anchors.topMargin: 0
        anchors.bottom: parent.bottom
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.left: parent.left
        anchors.leftMargin: 0
        currentIndex: tabBar.currentIndex
        Item {
            id: homeTab

            RowLayout {
                x: 174
                y: 80
                spacing: 60

                GridLayout {
                    rows: 1
                    columns: 2

                    Text {
                        id: element
                        text: qsTr("10")
                        Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                        verticalAlignment: Text.AlignTop
                        font.pixelSize: 48
                    }

                    Text {
                        id: element1
                        text: qsTr("BPM")
                        Layout.bottomMargin: 10
                        Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                        font.pixelSize: 20
                    }

                    Text {
                        id: element2
                        color: "#636060"
                        text: qsTr("BREATH RATE")
                        font.bold: false
                        Layout.columnSpan: 2
                        font.pixelSize: 16
                    }
                }

                GridLayout {
                    columns: 2
                    Text {
                        id: element3
                        text: qsTr("13")
                        verticalAlignment: Text.AlignTop
                        Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                        font.pixelSize: 48
                    }

                    Text {
                        id: element4
                        text: qsTr("cmH2O")
                        Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                        Layout.bottomMargin: 10
                        font.pixelSize: 20
                    }

                    Text {
                        id: element5
                        color: "#636060"
                        text: qsTr("PIP")
                        Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                        font.bold: false
                        font.pixelSize: 16
                        Layout.columnSpan: 2
                    }
                    rows: 1
                }

                GridLayout {
                    columns: 2
                    Text {
                        id: element6
                        text: qsTr("2.3")
                        verticalAlignment: Text.AlignTop
                        Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                        font.pixelSize: 48
                    }

                    Text {
                        id: element7
                        x: 84
                        text: qsTr("L")
                        verticalAlignment: Text.AlignTop
                        Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                        Layout.bottomMargin: 10
                        font.pixelSize: 20
                    }

                    Text {
                        id: element8
                        color: "#636060"
                        text: qsTr("MINUTE VOLUME")
                        Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                        font.bold: false
                        font.pixelSize: 16
                        Layout.columnSpan: 2
                    }
                    rows: 1
                }
            }

        }
        Item {
            id: discoverTab
        }
        Item {
            id: activityTab
        }
    }
}



