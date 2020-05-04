import QtQuick 2.0
import "./src/variables/fontawesome.js" as FontAwesome
import "./src/lists"
import "./config.js" as Config

Item {
    id: sidebar
    width: Config.sidebar_height
    property int currentView: 0
    Rectangle{

        id: sidebarrectangle
        width: 150
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
                    anchors.fill: parent
                    onItemClicked: {
                        if (menulist.model.get(index).text === "Monitor") {
                            console.log(0)
                            currentView = 0
                            menulist.model.get(0).class_name="light"
                            menulist.model.get(1).class_name="dark"
                            menulist.model.get(2).class_name="dark"
                            menulist.model.get(3).class_name="dark"
                        } else if (menulist.model.get(index).text === "Set Mode") {
                            console.log(1)
                            currentView = 1
                            menulist.model.get(0).class_name="dark"
                            menulist.model.get(1).class_name="light"
                            menulist.model.get(2).class_name="dark"
                            menulist.model.get(3).class_name="dark"
                        } else if(menulist.model.get(index).text === "Patient"){
                            console.log(2)
                            currentView = 2
                            menulist.model.get(0).class_name="dark"
                            menulist.model.get(1).class_name="dark"
                            menulist.model.get(2).class_name="light"
                            menulist.model.get(3).class_name="dark"
                        } else {
                            console.log(3)
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
            id: element2
            x: 88
            y: 186
            color: "#ffffff"
            text: qsTr("7.5")
            font.pixelSize: 38
        }

        Text {
            id: element3
            x: 68
            y: 227
            color: "#ffffff"
            text: qsTr("ExpMinVol")
            styleColor: "#ffffff"
            font.pixelSize: 16
        }

        Text {
            id: element4
            x: 9
            y: 202
            color: "#ffffff"
            text: qsTr("10")
            font.pixelSize: 12
        }

        Text {
            id: element5
            x: 13
            y: 239
            color: "#ffffff"
            text: qsTr("4")
            font.pixelSize: 12
        }

        Text {
            id: element6
            x: 114
            y: 246
            color: "#f9f9f9"
            text: qsTr("l/min")
            font.pixelSize: 12
        }

        Text {
            id: element7
            x: 79
            y: 275
            color: "#ffffff"
            text: qsTr("500")
            font.pixelSize: 38
        }

        Text {
            id: element8
            x: 114
            y: 316
            color: "#ffffff"
            text: qsTr("VTE")
            font.pixelSize: 16
            styleColor: "#ffffff"
        }

        Text {
            id: element9
            x: 8
            y: 291
            color: "#ffffff"
            text: qsTr("750")
            font.pixelSize: 12
        }

        Text {
            id: element10
            x: 8
            y: 321
            color: "#ffffff"
            text: qsTr("250")
            font.pixelSize: 12
        }

        Text {
            id: element11
            x: 129
            y: 334
            color: "#f9f9f9"
            text: qsTr("ml")
            font.pixelSize: 12
        }

        Text {
            id: element12
            x: 100
            y: 354
            color: "#ffffff"
            text: qsTr("20")
            horizontalAlignment: Text.AlignRight
            font.pixelSize: 38
        }

        Text {
            id: element13
            x: 111
            y: 395
            color: "#ffffff"
            text: qsTr("Rate")
            font.pixelSize: 16
            styleColor: "#ffffff"
        }

        Text {
            id: element14
            x: 8
            y: 370
            color: "#ffffff"
            text: qsTr("23")
            font.pixelSize: 12
        }

        Text {
            id: element15
            x: 12
            y: 400
            color: "#ffffff"
            text: qsTr("8")
            font.pixelSize: 12
        }

        Text {
            id: element16
            x: 110
            y: 413
            color: "#f9f9f9"
            text: qsTr("b/min")
            font.pixelSize: 12
        }

        Text {
            id: element17
            x: 13
            y: 438
            color: "#ffffff"
            text: "\uf015"
            font.pixelSize: 12
            font.family: webFont.name
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
    }

}
