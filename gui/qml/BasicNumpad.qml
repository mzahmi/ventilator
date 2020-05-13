import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0
import "./config.js" as Config

Item {
    id: name


    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.fill: parent

        Rectangle {
            id: numpadBG
            x: 420
            y: 125
            width: 212
            height: 266
            color: Config.bg_color
            anchors.right: parent.right
            anchors.left: flickable.right

            Row {
                id: element
                anchors.fill: parent

                Column {
                    id: column1
                    anchors.top: parent.top
                    anchors.topMargin: 0
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 0
                    anchors.left: parent.left
                    anchors.leftMargin: 0
                    width: parent.width/3

                    Button {
                        id: button1
                        anchors.right: parent.right
                        anchors.left: parent.left
                        anchors.leftMargin: 0
                        height: parent.height/4
                        onClicked: chosenText.text = chosenText.text+1
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {
                            font.pointSize: 20
                            text: qsTr("1")
                            anchors.fill: parent
                            //opacity: enabled ? 1.0 : 0.3
                            //color: control.down ? "#17a81a" : "#21be2b"
                            horizontalAlignment: Text.AlignHCenter
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                    }

                    Button {
                        id: button4
                        anchors.right: parent.right
                        anchors.left: parent.left
                        anchors.leftMargin: 0
                        height: parent.height/4
                        onClicked: chosenText.text = chosenText.text+4
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {

                            font.pointSize: 20
                            text: qsTr("4")
                            anchors.fill: parent
                            //opacity: enabled ? 1.0 : 0.3
                            //color: control.down ? "#17a81a" : "#21be2b"
                            horizontalAlignment: Text.AlignHCenter
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                    }

                    Button {
                        id: button7
                        onClicked: chosenText.text = chosenText.text+7
                        anchors.left: parent.left
                        anchors.leftMargin: 0
                        height: parent.height/4
                        anchors.right: parent.right
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {

                            font.pointSize: 20
                            text: qsTr("7")
                            anchors.fill: parent
                            //opacity: enabled ? 1.0 : 0.3
                            //color: control.down ? "#17a81a" : "#21be2b"
                            horizontalAlignment: Text.AlignHCenter
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                    }

                    Button {
                        id: button0
                        anchors.right: parent.right
                        anchors.left: parent.left
                        anchors.leftMargin: 0
                        height: parent.height/4
                        onClicked: chosenText.text = chosenText.text+"0"
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {
                            text: qsTr("0")
                            verticalAlignment: Text.AlignVCenter
                            anchors.fill: parent
                            horizontalAlignment: Text.AlignHCenter
                            font.pointSize: 20
                            elide: Text.ElideRight
                        }
                    }


                }

                Column {
                    id: column2
                    anchors.top: parent.top
                    anchors.topMargin: 0
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 0
                    anchors.left: column1.right
                    width: parent.width/3

                    Button {
                        id: button2
                        anchors.right: parent.right
                        anchors.left: parent.left
                        height: parent.height/4
                        onClicked: chosenText.text = chosenText.text+2
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {
                            font.pointSize: 20
                            text: qsTr("2")
                            anchors.fill: parent
                            //opacity: enabled ? 1.0 : 0.3
                            //color: control.down ? "#17a81a" : "#21be2b"
                            horizontalAlignment: Text.AlignHCenter
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                    }

                    Button {
                        id: button5
                        anchors.right: parent.right
                        anchors.left: parent.left
                        height: parent.height/4
                        onClicked: chosenText.text = chosenText.text+5
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {

                            font.pointSize: 20
                            text: qsTr("5")
                            anchors.fill: parent
                            //opacity: enabled ? 1.0 : 0.3
                            //color: control.down ? "#17a81a" : "#21be2b"
                            horizontalAlignment: Text.AlignHCenter
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }

                    }

                    Button {
                        id: button8
                        onClicked: chosenText.text = chosenText.text+8
                        anchors.right: parent.right
                        anchors.left: parent.left
                        height: parent.height/4
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {

                            font.pointSize: 20
                            text: qsTr("8")
                            anchors.fill: parent
                            //opacity: enabled ? 1.0 : 0.3
                            //color: control.down ? "#17a81a" : "#21be2b"
                            horizontalAlignment: Text.AlignHCenter
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                    }

                    Button {
                        id: buttonDot
                        anchors.right: parent.right
                        anchors.left: parent.left
                        anchors.leftMargin: 0
                        height: parent.height/4
                        onClicked: chosenText.text = chosenText.text+"."
                        palette {
                            button: Constants.mainbg
                        }
                        contentItem: Text {
                            text: qsTr(".")
                            verticalAlignment: Text.AlignVCenter
                            anchors.fill: parent
                            horizontalAlignment: Text.AlignHCenter
                            font.pointSize: 20
                            elide: Text.ElideRight
                        }
                    }


                }
                Column {
                    id: column3
                    anchors.top: parent.top
                    anchors.topMargin: 0
                    anchors.bottom: parent.bottom
                    anchors.bottomMargin: 0
                    anchors.left: column2.right
                    width: parent.width/3

                    Button {
                        id: button9
                        anchors.right: parent.right
                        anchors.leftMargin: 0
                        anchors.left: parent.left
                        height: parent.height/4
                        contentItem: Text {
                            text: qsTr("9")
                            anchors.fill: parent
                            horizontalAlignment: Text.AlignHCenter
                            font.pointSize: 20
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                    }

                    Button {
                        id: button6
                        anchors.right: parent.right
                        anchors.left: parent.left
                        height: parent.height/4
                        contentItem: Text {
                            text: qsTr("6")
                            anchors.fill: parent
                            horizontalAlignment: Text.AlignHCenter
                            font.pointSize: 20
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                        anchors.leftMargin: 0
                    }

                    Button {
                        id: button3
                        anchors.right: parent.right
                        anchors.left: parent.left
                        height: parent.height/4
                        contentItem: Text {
                            text: qsTr("3")
                            anchors.fill: parent
                            horizontalAlignment: Text.AlignHCenter
                            font.pointSize: 20
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                        anchors.leftMargin: 0
                    }

                    Button {
                        id: buttonDelete
                        anchors.right: parent.right
                        anchors.left: parent.left
                        height: parent.height/4
                        anchors.leftMargin: 0
                        contentItem: Text {
                            text: qsTr("<")
                            anchors.fill: parent
                            horizontalAlignment: Text.AlignHCenter
                            font.pointSize: 20
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                    }
                }



            }

        }
    }
}


/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:9;anchors_width:22.222222222222225}D{i:18;anchors_width:22.222222222222225}
D{i:23;anchors_width:22.222222222222225}D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/
