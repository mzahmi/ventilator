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
        Item{
            id: infoStack
            height: 240
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            Layout.fillWidth: true



            TabBar {
                id: tabBar
                currentIndex: 0
                anchors.right: parent.right
                anchors.rightMargin: 0
                anchors.left: parent.left
                anchors.leftMargin: 0
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
            }

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
                        y: 79
                        transformOrigin: Item.Center
                        anchors.left: parent.left
                        anchors.leftMargin: 197
                        anchors.right: parent.right
                        anchors.rightMargin: 197
                        spacing: 30

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




    }
}

/*##^##
Designer {
    D{i:6;anchors_width:800}D{i:17;anchors_x:174;anchors_y:80}D{i:15;anchors_width:800}
D{i:2;anchors_x:0;anchors_y:0}
}
##^##*/
