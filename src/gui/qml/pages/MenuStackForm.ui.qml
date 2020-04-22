import QtQuick 2.4
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.13

Item {
    id: element9
    width: 800
    height: 400


    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.top: parent.top
        anchors.bottom: parent.bottom
        anchors.left: parent.left
        anchors.leftMargin: 0
        anchors.right: parent.right

        StackLayout {
            anchors.left: parent.left
            anchors.leftMargin: 0
            anchors.topMargin: 0
            anchors.rightMargin: 0
            Item {
                id: homeTab
                RowLayout {
                    y: 79
                    anchors.left: parent.left
                    anchors.leftMargin: 197
                    anchors.rightMargin: 197
                    GridLayout {
                        Text {
                            id: element
                            text: qsTr("10")
                            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                            font.pixelSize: 48
                            verticalAlignment: Text.AlignTop
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
                        columns: 2
                        rows: 1
                    }

                    GridLayout {
                        Text {
                            id: element3
                            text: qsTr("13")
                            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                            font.pixelSize: 48
                            verticalAlignment: Text.AlignTop
                        }

                        Text {
                            id: element4
                            text: qsTr("cmH2O")
                            Layout.bottomMargin: 10
                            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                            font.pixelSize: 20
                        }

                        Text {
                            id: element5
                            color: "#636060"
                            text: qsTr("PIP")
                            font.bold: false
                            Layout.columnSpan: 2
                            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                            font.pixelSize: 16
                        }
                        rows: 1
                        columns: 2
                    }

                    GridLayout {
                        Text {
                            id: element6
                            text: qsTr("2.3")
                            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                            font.pixelSize: 48
                            verticalAlignment: Text.AlignTop
                        }

                        Text {
                            id: element7
                            x: 84
                            text: qsTr("L")
                            Layout.bottomMargin: 10
                            Layout.alignment: Qt.AlignLeft | Qt.AlignBottom
                            font.pixelSize: 20
                            verticalAlignment: Text.AlignTop
                        }

                        Text {
                            id: element8
                            color: "#636060"
                            text: qsTr("MINUTE VOLUME")
                            font.bold: false
                            Layout.columnSpan: 2
                            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
                            font.pixelSize: 16
                        }
                        rows: 1
                        columns: 2
                    }
                    spacing: 30
                    anchors.right: parent.right
                    transformOrigin: Item.Center
                }
            }

            Item {
                id: discoverTab
            }

            Item {
                id: activityTab
            }
            anchors.bottom: parent.bottom
            anchors.top: parent.top
            currentIndex: tabBar.currentIndex
            anchors.right: parent.right
        }

        TabBar {
            id: tabBar
            anchors.left: parent.left
            TabButton {
                text: qsTr("Home")
                background: Rectangle {
                    color: tabBar.currentIndex === 0 ? "#84a5c3" : "#edf0f4"
                }
            }

            TabButton {
                text: qsTr("Therapy")
                background: Rectangle {
                    color: tabBar.currentIndex === 1 ? "#84a5c3" : "#edf0f4"
                }
            }

            TabButton {
                text: qsTr("Monitors")
                background: Rectangle {
                    color: tabBar.currentIndex === 2 ? "#84a5c3" : "#edf0f4"
                }
            }

            TabButton {
                text: qsTr("Menu")
                background: Rectangle {
                    color: tabBar.currentIndex === 3 ? "#84a5c3" : "#edf0f4"
                }
            }
            anchors.right: parent.right
            currentIndex: 0
        }

    }
}/*##^##
    Designer {
        D{i:1;anchors_width:800}D{i:12;anchors_x:174;anchors_y:80}D{i:10;anchors_width:800}
    }
    ##^##*/

