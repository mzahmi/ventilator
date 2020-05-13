import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import "./src/variables/fontawesome.js"
as FontAwesome
import "./src/lists"
import "./material/qml/material"
import "./config.js"
as Config
import "."

Item {
    id: root
    width: 800
    height: 480
    signal alarm(string status)
    Component.onCompleted: {
        AlarmManager.alarmStatus.connect(root.alarm)
    }
    onAlarm:{
        if (status !== "none"){
            alarm.visible=true
        } else{
            alarm.visible = false
        }
        status1.text = AlarmManager.status
        title1.text = AlarmManager.title
        info1.text = AlarmManager.info
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


                ViewMonitor {
                    anchors.fill: parent
                    onPresetClicked: {
                        sidebar.openTab()
                        viewmodeview.presetClicked()

                    }
                }
            }

            Item {
                id: viewmode
                ViewMode {
                    id: viewmodeview
                    anchors.fill: parent
                }
            }
            Item {
                id: viewpatient
                ViewPatient {

                }

            }
            Item {
                id: viewsettings
                ViewSettings {

                }

            }
        }
    }

    FontLoader {
        id: webFont;source: "./src/variables/fontawesome-webfont.ttf"
    }

    Rectangle {
        id: alarm
        x: 327
        y: 127
        width: 200
        height: 200
        color: "#d10000"
        visible: false

        RowLayout {
            x: 2
            y: 3

            ColumnLayout {
            }

            ColumnLayout {
            }
        }

        ColumnLayout {
            x: 13
            y: 12

            RowLayout {

                Text {
                    id: status1
                    color: "#ffffff"
                    text: qsTr("Text")
                    font.pixelSize: 18
                }
            }

            RowLayout {

                Text {
                    id: title1
                    color: "#ffffff"
                    text: qsTr("Text")
                    font.pixelSize: 18
                }
            }

            RowLayout {

                Text {
                    id: info
                    color: "#ffffff"
                    text: qsTr("Info")
                    font.pixelSize: 12
                }

                Text {
                    id: info1
                    color: "#ffffff"
                    text: qsTr("Text")
                    font.pixelSize: 12
                }
            }
        }
    }

    SideBar {
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
