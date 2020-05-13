import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import "./src/variables/fontawesome.js" as FontAwesome
import "./src/lists"
import "./material/qml/material"
import "./config.js" as Config
import "."

Item {
    id: element
    width: 800
    height: 480
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


                ViewMonitor{
                    anchors.fill: parent
                    onPresetClicked: {
                        sidebar.openTab()
                        viewmodeview.presetClicked()

                    }
                }
            }

            Item {
                id: viewmode
                ViewMode{
                    id: viewmodeview
                    anchors.fill: parent
                }
            }
            Item {
                id: viewpatient
                ViewPatient{

                }

            }
            Item {
                id: viewsettings
                ViewSettings{

                }

            }
        }
    }

    FontLoader { id: webFont; source: "./src/variables/fontawesome-webfont.ttf" }

    SideBar{
        id: sidebar
        anchors.left: parent.left
        anchors.leftMargin: 0
        anchors.bottom: parent.bottom
        anchors.top: parent.top
        anchors.topMargin: 0
        width: Config.sidebar_width

    }


}

/*##^##
Designer {
    D{i:1;anchors_width:150}D{i:12;anchors_height:400;anchors_width:200}
}
##^##*/
