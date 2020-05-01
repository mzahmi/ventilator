import QtQuick 2.4
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0

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

                TabHome {
                    id: tabHome
                    anchors.top: parent.top
                    anchors.topMargin: 0
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 0
                }


            }

            Item {
                id: therapyTab

                TabTherapy {
                    id: tabTherapy
                    anchors.top: parent.top
                    anchors.topMargin: 0
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 0
                }
            }

            Item {
                id: monitorsTab
            }
            Item {
                id: menuTab
            }
            anchors.bottom: parent.bottom
            anchors.top: parent.top
            currentIndex: tabBar.currentIndex
            anchors.right: parent.right
        }

        TabBar {
            id: tabBar
            height: 40
            anchors.left: parent.left
            TabButton {
                anchors.bottom: parent.bottom
                anchors.top: parent.top
                anchors.topMargin: 0
                background: Rectangle {
                    color: tabBar.currentIndex === 0 ? "#84a5c3" : "#edf0f4"
                }
                contentItem: Text {
                    text: "Home"
                    color: tabBar.currentIndex === 0 ? "white" : "black"
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }
            }

            TabButton {
                anchors.bottom: parent.bottom
                anchors.top: parent.top
                anchors.topMargin: 0
                background: Rectangle {
                    color: tabBar.currentIndex === 1 ? "#84a5c3" : "#edf0f4"
                }
                contentItem: Text {
                    text: "Therapy"
                    color: tabBar.currentIndex === 1 ? "white" : "black"
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }
            }

            TabButton {
                anchors.bottom: parent.bottom
                anchors.top: parent.top
                anchors.topMargin: 0
                background: Rectangle {
                    color: tabBar.currentIndex === 2 ? "#84a5c3" : "#edf0f4"
                    MouseArea{
                        onClicked: ButtonBridge.sendToGo("hello from qml")
                    }
                }
                contentItem: Text {
                    text: "Monitors"
                    color: tabBar.currentIndex === 2 ? "white" : "black"
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }
            }

            TabButton {
                anchors.bottom: parent.bottom
                anchors.top: parent.top
                anchors.topMargin: 0
                background: Rectangle {
                    color: tabBar.currentIndex === 3 ? "#84a5c3" : "#edf0f4"
                }
                contentItem: Text {
                    text: "Menu"
                    color: tabBar.currentIndex === 3 ? "white" : "black"
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }
            }
            anchors.right: parent.right
            currentIndex: 1
        }

    }
}/*##^##
Designer {
    D{i:4;anchors_y:160}D{i:12;anchors_x:174;anchors_y:80}D{i:10;anchors_width:800}D{i:1;anchors_width:800}
}
##^##*/
