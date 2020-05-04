import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.3
import "./src/variables/fontawesome.js" as FontAwesome
import "./src/lists"
import "./material/qml/material"


Item {
    id: element
    width: 800
    height: 460
    FontLoader { id: webFont; source: "./src/variables/fontawesome-webfont.ttf" }

    Rectangle{


        id: sidebar
        width: 150
        color: "#444444"
        anchors.bottom: parent.bottom
        anchors.left: parent.left
        anchors.top: parent.top

        Column {
            id: column
            anchors.fill: parent

            Item {
                id: element1
                height: 170
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0
                IconListView{
                    anchors.bottom: parent.bottom
                    anchors.right: parent.right
                    anchors.left: parent.left
                    anchors.top: parent.top
                    model: [{
                            text: "Monitor",
                            leftIcon: "\uf06e"
                        },
                        {
                            text: "Set Mode",
                            leftIcon: "\uf067",
                            class_name: "dark"
                        },
                        {
                            text: "Patient",
                            leftIcon: "\uf007",
                            class_name: "dark"

                        },
                        {
                            text: "Settings",
                            leftIcon: "\uf044",
                            class_name: "dark"

                        }]

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

            Item {
                id: mainwindow

                anchors.fill: parent
                MonitorView{
                    anchors.fill: parent


                }
            }

            Item {
                id: setmodewindow
            }
        }
    }

}

/*##^##
Designer {
    D{i:3;anchors_width:150}D{i:2;anchors_height:400;anchors_width:200}D{i:20;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46}
D{i:21;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46}D{i:22;anchors_height:200;anchors_width:200;anchors_x:12;anchors_y:108}
D{i:26;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46}D{i:27;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46}
D{i:25;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46}D{i:24;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46}
D{i:23;anchors_height:200;anchors_width:200;anchors_x:18;anchors_y:46}
}
##^##*/
