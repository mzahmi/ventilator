import QtQuick 2.0
import "./src/variables/fontawesome.js" as FontAwesome
import "./src/lists"
import "./config.js" as Config
import QtQuick.Layouts 1.0
import "."

Item {
    id: sidebar
    width: Config.sidebar_height
    property int currentView: 0
    // patient information when first started
    Component.onCompleted:{
        currentView = 2
        menulist.model.get(0).class_name="dark"
        menulist.model.get(1).class_name="dark"
        menulist.model.get(2).class_name="light"
        menulist.model.get(3).class_name="dark"
        ModeSelect.modeSelected.connect(sidebar.openMonitor)
    }
    signal openTab()
    onOpenTab: {
        currentView = 1
        menulist.model.get(0).class_name="dark"
        menulist.model.get(1).class_name="light"
        menulist.model.get(2).class_name="dark"
        menulist.model.get(3).class_name="dark"
    }
    signal openMonitor()
    onOpenMonitor: {
        currentView = 0
        menulist.model.get(0).class_name="light"
        menulist.model.get(1).class_name="dark"
        menulist.model.get(2).class_name="dark"
        menulist.model.get(3).class_name="dark"
    }

    Rectangle{
        id: sidebarrectangle
        width: Config.sidebar_width
        color: Config.bg_color
        anchors.bottom: parent.bottom
        anchors.left: parent.left
        anchors.top: parent.top

        Column {
            id: column
            anchors.fill: parent

            Item {
                id: menuitem
                height: 170
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0

                IconListView{
                    id: menulist
                    anchors.rightMargin: 1
                    anchors.fill: parent

                    onItemClicked: {
                        if (menulist.model.get(index).text === "Monitor") {
                            currentView = 0
                            menulist.model.get(0).class_name="light"
                            menulist.model.get(1).class_name="dark"
                            menulist.model.get(2).class_name="dark"
                            menulist.model.get(3).class_name="dark"
                        } else if (menulist.model.get(index).text === "Set Mode") {
                            currentView = 1
                            menulist.model.get(0).class_name="dark"
                            menulist.model.get(1).class_name="light"
                            menulist.model.get(2).class_name="dark"
                            menulist.model.get(3).class_name="dark"
                        } else if(menulist.model.get(index).text === "Patient"){
                            currentView = 2
                            menulist.model.get(0).class_name="dark"
                            menulist.model.get(1).class_name="dark"
                            menulist.model.get(2).class_name="light"
                            menulist.model.get(3).class_name="dark"
                        } else {
                            currentView = 3
                            menulist.model.get(0).class_name="dark"
                            menulist.model.get(1).class_name="dark"
                            menulist.model.get(2).class_name="dark"
                            menulist.model.get(3).class_name="light"
                        }
                    }
                    model: ListModel{

                        ListElement{
                            text: "Monitor"
                            leftIcon: "\uf06e"
                            class_name: "light"
                        }

                        ListElement{
                            text: "Set Mode"
                            divider: "Divider 1"
                            leftIcon: "\uf067"
                            class_name: "dark"
                        }
                        ListElement{
                            text: "Patient"
                            divider: "Divider 1"
                            leftIcon: "\uf007"
                            class_name: "dark"
                        }
                        ListElement{
                            text: "Settings"
                            divider: "Divider 1"
                            leftIcon: "\uf044"
                            class_name: "dark"
                        }
                    }
                }
            }
        }

        Text {
            id: element17
            x: 13
            y: 438
            color: "#ffffff"
            text: "\uf015"
            font.pixelSize: 12
            font.family: webFont.name

            MouseArea {
                id: fullscreentoggle
                x: -10
                width: 38
                height: 26
                onClicked: {
                   var fs=false
                }
            }
        }

        Text {
            id: element18
            x: 68
            y: 438
            color: "#ffffff"
            text: "\uf06a"
            font.pixelSize: 12
        }

        Text {
            id: element19
            x: 122
            y: 438
            color: "#ffffff"
            text: "\uf043"
            font.pixelSize: 12
        }

        Row {
            id: monitorrow
            property int pip:25
            property int vt: 10
            property int rate: 15
            property int peep: 4
            property int fio2: 25
            property string mode: "PAC"

            y: 170
            spacing: 0
            anchors.left: parent.left
            anchors.right: parent.right



            Column {
                id: column1
                width:Config.sidebar_width/2
                Layout.fillHeight: false
                Layout.fillWidth: false
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter

                Text {
                    id: piptext
                    color: "#ffffff"
                    text: monitorrow.pip
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    anchors.right: parent.right
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                    font.pixelSize: Config.sidebar_livetext_size1
                }

                Text {
                    id: element2
                    color: "#ffffff"
                    text: qsTr("PIP")
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    anchors.right: parent.right
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                    styleColor: "#ffffff"
                    font.pixelSize: Config.sidebar_livetext_size2
                }

                Text {
                    id: element3
                    color: "#f9f9f9"
                    text: qsTr("cmH2O")
                    anchors.right: parent.right
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                    font.pixelSize: Config.sidebar_livetext_size3
                }
            }

            Column {
                id: column2
                width: Config.sidebar_width/2
                Layout.fillWidth: false
                Text {
                    id: element4
                    color: "#ffffff"
                    text: monitorrow.vt
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    anchors.right: parent.right
                    font.pixelSize: Config.sidebar_livetext_size1
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                }

                Text {
                    id: element5
                    color: "#ffffff"
                    text: qsTr("Vt")
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    anchors.right: parent.right
                    font.pixelSize: Config.sidebar_livetext_size2
                    styleColor: "#ffffff"
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                }

                Text {
                    id: element7
                    color: "#f9f9f9"
                    text: qsTr("ml")
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    anchors.right: parent.right
                    font.pixelSize: Config.sidebar_livetext_size3
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }

        RowLayout {
            y: 257
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
                width: Config.sidebar_height/2
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

        RowLayout {
            y: 345
            spacing: 0
            anchors.left: parent.left
            anchors.right: parent.right
            ColumnLayout {
                width: Config.sidebar_height/2
                id: column5
                Text {
                    id: element14
                    color: "#ffffff"
                    text: monitorrow.fio2
                    font.pixelSize: Config.sidebar_livetext_size1
                    Layout.fillWidth: true
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }

                Text {
                    id: element15
                    color: "#ffffff"
                    text: qsTr("FIO2")
                    font.pixelSize: Config.sidebar_livetext_size2
                    styleColor: "#ffffff"
                    Layout.fillWidth: true
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }

                Text {
                    id: element16
                    color: "#f9f9f9"
                    text: qsTr("%")
                    font.pixelSize: Config.sidebar_livetext_size3
                    Layout.fillWidth: true
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }
            }

            ColumnLayout {
                width: Config.sidebar_height/2
                id: column6
                Text {
                    id: element20
                    color: "#ffffff"
                    text: monitorrow.mode
                    font.pixelSize: Config.sidebar_livetext_size1
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                }

                Text {
                    id: element21
                    color: "#ffffff"
                    text: qsTr("Mode")
                    font.pixelSize: Config.sidebar_livetext_size2
                    styleColor: "#ffffff"
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                }

                Text {
                    id: element22
                    color: "#f9f9f9"
                    font.pixelSize: Config.sidebar_livetext_size3
                    horizontalAlignment: Text.AlignHCenter
                    Layout.fillWidth: true
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }
    }

}

/*##^##
Designer {
    D{i:0;formeditorZoom:1.25}D{i:14;anchors_x:2}D{i:23;anchors_x:2}D{i:32;anchors_x:2}
}
##^##*/
