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
    id: sidebar
    width: Config.sidebar_height
    property int currentView: 0
    // patient information when first started
    Component.onCompleted: {
        currentView = 2
        menulist.model.get(0).class_name = "dark"
        menulist.model.get(1).class_name = "dark"
        menulist.model.get(2).class_name = "light"
        menulist.model.get(3).class_name = "dark"
        ModeSelect.modeSelected.connect(sidebar.openMonitor)
    }
    signal openTab()
    onOpenTab: {
        currentView = 1
        menulist.model.get(0).class_name = "dark"
        menulist.model.get(1).class_name = "light"
        menulist.model.get(2).class_name = "dark"
        menulist.model.get(3).class_name = "dark"
    }
    signal openMonitor()
    onOpenMonitor: {
        currentView = 0
        menulist.model.get(0).class_name = "light"
        menulist.model.get(1).class_name = "dark"
        menulist.model.get(2).class_name = "dark"
        menulist.model.get(3).class_name = "dark"
    }

    Rectangle {
        id: sidebarrectangle
        width: Config.sidebar_width
        color: Config.bg_color
        anchors.bottom: parent.bottom
        anchors.left: parent.left
        anchors.top: parent.top

        Column {
            id: menulistcolumn
            anchors.right: parent.right
            anchors.left: parent.left
            anchors.top: parent.top

            Item {
                id: menuitem
                height: 170
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0

                IconListView {
                    id: menulist
                    anchors.rightMargin: 1
                    anchors.fill: parent

                    onItemClicked: {
                        if (menulist.model.get(index).text === "Monitor") {
                            currentView = 0
                            menulist.model.get(0).class_name = "light"
                            menulist.model.get(1).class_name = "dark"
                            menulist.model.get(2).class_name = "dark"
                            menulist.model.get(3).class_name = "dark"
                        } else if (menulist.model.get(index).text === "Set Mode") {
                            currentView = 1
                            menulist.model.get(0).class_name = "dark"
                            menulist.model.get(1).class_name = "light"
                            menulist.model.get(2).class_name = "dark"
                            menulist.model.get(3).class_name = "dark"
                        } else if (menulist.model.get(index).text === "Patient") {
                            currentView = 2
                            menulist.model.get(0).class_name = "dark"
                            menulist.model.get(1).class_name = "dark"
                            menulist.model.get(2).class_name = "light"
                            menulist.model.get(3).class_name = "dark"
                        } else {
                            currentView = 3
                            menulist.model.get(0).class_name = "dark"
                            menulist.model.get(1).class_name = "dark"
                            menulist.model.get(2).class_name = "dark"
                            menulist.model.get(3).class_name = "light"
                        }
                    }
                    model: ListModel {

                        ListElement {
                            text: "Monitor"
                            leftIcon: "\uf06e"
                            class_name: "light"
                        }

                        ListElement {
                            text: "Set Mode"
                            divider: "Divider 1"
                            leftIcon: "\uf067"
                            class_name: "dark"
                        }
                        ListElement {
                            text: "Patient"
                            divider: "Divider 1"
                            leftIcon: "\uf007"
                            class_name: "dark"
                        }
                        ListElement {
                            text: "Settings"
                            divider: "Divider 1"
                            leftIcon: "\uf044"
                            class_name: "dark"
                        }
                    }
                }
            }
        }


        SwipeView {
            id: view
            anchors.bottom: iconrow.top
            anchors.top: menulistcolumn.bottom
            anchors.right: parent.right
            anchors.left: parent.left

            currentIndex: 0

            Item {
                id: firstPage
                LiveData {
                    id: liveData
                    anchors.bottom: iconrow.top
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    anchors.right: parent.right
                    anchors.top: menulistcolumn.bottom
                }
            }
            Item {
                id: secondPage
                LiveSetData{
                    id: liveSetData
                    anchors.bottom: iconrow.top
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    anchors.right: parent.right
                    anchors.top: menulistcolumn.bottom
                }
            }
        }


        RowLayout {
            id: iconrow
            y: 438
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 20
            anchors.leftMargin: 15
            anchors.right: parent.right
            anchors.left: parent.left

            Text {
                id: element17
                color: "#ffffff"
                text: "\uf015"
                font.pixelSize: 12
                font.family: webFont.name
            }

            Text {
                id: element18
                color: "#ffffff"
                text: "\uf06a"
                font.pixelSize: 12
            }

            Text {
                id: element19
                color: "#ffffff"
                text: "\uf043"
                font.pixelSize: 12
            }
        }
    }

}

/*##^##
Designer {
    D{i:0;formeditorZoom:1.25}
}
##^##*/
