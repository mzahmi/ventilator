import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.3
import "./src/variables/fontawesome.js" as FontAwesome
import "./src/lists"
import "./material/qml/material"
import "./config.js" as Config

//        Button {
//            id: button
//            x: 34
//            y: 218
//            text: qsTr("Button")
//            Connections{
//                target: QmlBridge
//            }
//            onClicked: QmlBridge.sendToGo("hello from qml")
//        }


Item {
    id: element
    width: 800
    height: 460
    FontLoader { id: webFont; source: "./src/variables/fontawesome-webfont.ttf" }

    Menulist{
        id: sidebar
        anchors.left: parent.left
        anchors.leftMargin: 0
        anchors.bottom: parent.bottom
        anchors.top: parent.top
        anchors.topMargin: 0
        width: Config.sidebar_width

    }

    Rectangle {
        id: mainview
        color: "#ffffff"
        anchors.left: sidebar.right
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        anchors.top: parent.top


        StackLayout {
            id: stackLayout
            anchors.fill: parent
            currentIndex: sidebar.currentView

            Item {
                id: viewmonitor
                anchors.fill: parent
                ViewMonitor{
                    anchors.fill: parent
                }
            }

            Item {
                id: viewmode
                anchors.fill: parent
                ViewMode{
                    anchors.fill: parent
                }
            }
            Item {
                id: viewpatient
                anchors.fill: parent
                ViewPatient{

                }

            }
            Item {
                id: viewsettings
                anchors.fill: parent
                ViewSettings{

                }

            }
        }
    }

}

/*##^##
Designer {
    D{i:2;anchors_height:400;anchors_width:200}D{i:3;anchors_width:150}
}
##^##*/
