import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.0
import "./config.js" as Config

Item {
    id: name
    height: 800


    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.fill: parent

        Flickable {
            id: flickable
            anchors.rightMargin: 0
            anchors.leftMargin: 0
            contentHeight: 650
            anchors.fill: parent

            Text {
                id: grouptitle
                x: 38
                y: 14
                text: qsTr("Pressure Assist/Control")
                font.pixelSize: 24
            }

            Text {
                id: ieratiotitle
                x: 140
                y: 59
                text: "I/E Ratio"
                font.pixelSize: 16
                font.family: "Open Sans"
            }

            Row {
                x: 173
                y: 88
                width: 294
                height: 40
                spacing: 50

                RadioButton {
                    id: control
                    text: qsTr("1:1")
                    ButtonGroup.group: radioGroup

                    indicator: Rectangle {
                        implicitWidth: 20
                        implicitHeight: 20
                        x: control.leftPadding
                        y: parent.height / 2 - height / 2
                        radius: 13
                        border.color: control.down ? Config.color_primary : Config.color_primary

                        Rectangle {
                            width: 8
                            height: 8
                            x: 6
                            y: 6
                            radius: 7
                            color: control.down ? Config.color_primary : Config.color_primary
                            visible: control.checked
                        }
                    }

                    contentItem: Text {
                        text: control.text
                        anchors.left: parent.left
                        anchors.leftMargin: 32
                        horizontalAlignment: Text.AlignLeft
                        font: control.font
                        opacity: enabled ? 1.0 : 0.3
                        verticalAlignment: Text.AlignVCenter

                    }
                }
                RadioButton {
                    id: control2
                    text: qsTr("1:1")
                    ButtonGroup.group: radioGroup
                    checked: true

                    indicator: Rectangle {
                        implicitWidth: 20
                        implicitHeight: 20
                        x: control2.leftPadding
                        y: parent.height / 2 - height / 2
                        radius: 13
                        border.color: control2.down ? Config.color_primary : Config.color_primary

                        Rectangle {
                            width: 8
                            height: 8
                            x: 6
                            y: 6
                            radius: 7
                            color: control2.down ? Config.color_primary : Config.color_primary
                            visible: control2.checked
                        }
                    }

                    contentItem: Text {
                        text: "1:2"
                        anchors.left: parent.left
                        anchors.leftMargin: 32
                        horizontalAlignment: Text.AlignLeft
                        font: control2.font
                        opacity: enabled ? 1.0 : 0.3
                        verticalAlignment: Text.AlignVCenter

                    }
                }

                RadioButton {
                    id: control3
                    text: qsTr("1:1")
                    ButtonGroup.group: radioGroup
                    indicator: Rectangle {
                        x: control3.leftPadding
                        y: parent.height / 2 - height / 2
                        radius: 13
                        Rectangle {
                            x: 6
                            y: 6
                            width: 8
                            height: 8
                            color: control3.down ? Config.color_primary : Config.color_primary
                            radius: 7
                            visible: control3.checked
                        }
                        implicitHeight: 20
                        border.color: control3.down ? Config.color_primary : Config.color_primary
                        implicitWidth: 20
                    }
                    contentItem: Text {
                        text: "1:3"
                        verticalAlignment: Text.AlignVCenter
                        anchors.left: parent.left
                        opacity: enabled ? 1.0 : 0.3
                        horizontalAlignment: Text.AlignLeft
                        font: control3.font
                        anchors.leftMargin: 32
                    }
                }

                RadioButton {
                    id: control4
                    text: qsTr("1:4")
                    ButtonGroup.group: radioGroup
                    indicator: Rectangle {
                        x: control4.leftPadding
                        y: parent.height / 2 - height / 2
                        radius: 13
                        Rectangle {
                            x: 6
                            y: 6
                            width: 8
                            height: 8
                            color: control4.down ? Config.color_primary : Config.color_primary
                            radius: 7
                            visible: control4.checked
                        }
                        implicitHeight: 20
                        border.color: control4.down ? Config.color_primary : Config.color_primary
                        implicitWidth: 20
                    }
                    contentItem: Text {
                        text: control4.text
                        verticalAlignment: Text.AlignVCenter
                        anchors.left: parent.left
                        opacity: enabled ? 1.0 : 0.3
                        horizontalAlignment: Text.AlignLeft
                        font: control4.font
                        anchors.leftMargin: 32
                    }
                }

            }




            Slider {
                id: slider2
                x: 136
                y: 160
                width: 369
                height: 30
                background: Rectangle {
                    x: slider2.leftPadding
                    y: slider2.topPadding + slider2.availableHeight / 2 - height / 2
                    width: slider2.availableWidth
                    height: implicitHeight
                    color: "#bdbebf"
                    radius: 2
                    Rectangle {
                        width: slider2.visualPosition * parent.width
                        height: parent.height
                        color: Config.color_primary
                        radius: 2
                    }
                    implicitHeight: 4
                    implicitWidth: 200
                }
                spacing: 0
                value: 25
                stepSize: 5
                Text {
                    id: title2
                    x: 5
                    y: -20
                    text: "PIP: "+slider2.value
                    font.pixelSize: 16
                    font.family: "Open Sans"
                }

                Text {
                    id: max2
                    x: 375
                    y: 7
                    text: slider2.to
                    font.pixelSize: 12
                }

                Text {
                    id: min2
                    x: -13
                    y: 7
                    text: slider2.from
                    font.pixelSize: 12
                }
                handle: Rectangle {
                    x: slider2.leftPadding + slider2.visualPosition * (slider2.availableWidth - width)
                    y: slider2.topPadding + slider2.availableHeight / 2 - height / 2
                    color: slider2.pressed ? "#f0f0f0" : "#f6f6f6"
                    radius: 13
                    implicitHeight: 18
                    border.color: "#bdbebf"
                    implicitWidth: 18
                }
                snapMode: "SnapAlways"
                from: 15
                wheelEnabled: false
                to: 40
            }

            Slider {
                id: slider3
                x: 137
                y: 222
                width: 369
                height: 30
                background: Rectangle {
                    x: slider3.leftPadding
                    y: slider3.topPadding + slider3.availableHeight / 2 - height / 2
                    width: slider3.availableWidth
                    height: implicitHeight
                    color: "#bdbebf"
                    radius: 2
                    Rectangle {
                        width: slider3.visualPosition * parent.width
                        height: parent.height
                        color: Config.color_primary
                        radius: 2
                    }
                    implicitHeight: 4
                    implicitWidth: 200
                }
                stepSize: 2
                value: 20
                spacing: 0
                Text {
                    id: title3
                    x: 5
                    y: -20
                    text: "BPM: "+slider3.value
                    font.pixelSize: 16
                    font.family: "Open Sans"
                }

                Text {
                    id: max3
                    x: 375
                    y: 7
                    text: slider3.to
                    font.pixelSize: 12
                }

                Text {
                    id: min3
                    x: -13
                    y: 7
                    text: slider3.from
                    font.pixelSize: 12
                }
                handle: Rectangle {
                    x: slider3.leftPadding + slider3.visualPosition * (slider3.availableWidth - width)
                    y: slider3.topPadding + slider3.availableHeight / 2 - height / 2
                    color: slider3.pressed ? "#f0f0f0" : "#f6f6f6"
                    radius: 13
                    implicitHeight: 18
                    border.color: "#bdbebf"
                    implicitWidth: 18
                }
                snapMode: "SnapAlways"
                from: 8
                wheelEnabled: false
                to: 40
            }

            Slider {
                id: slider4
                x: 137
                y: 284
                width: 369
                height: 30
                background: Rectangle {
                    x: slider4.leftPadding
                    y: slider4.topPadding + slider4.availableHeight / 2 - height / 2
                    width: slider4.availableWidth
                    height: implicitHeight
                    color: "#bdbebf"
                    radius: 2
                    Rectangle {
                        width: slider4.visualPosition * parent.width
                        height: parent.height
                        color: Config.color_primary
                        radius: 2
                    }
                    implicitHeight: 4
                    implicitWidth: 200
                }
                spacing: 0
                value: 20
                stepSize: 5
                Text {
                    id: title4
                    x: 5
                    y: -20
                    text: "PMAX: "+slider4.value
                    font.pixelSize: 16
                    font.family: "Open Sans"
                }

                Text {
                    id: max4
                    x: 375
                    y: 7
                    text: slider4.to
                    font.pixelSize: 12
                }

                Text {
                    id: min4
                    x: -13
                    y: 7
                    text: slider4.from
                    font.pixelSize: 12
                }
                handle: Rectangle {
                    x: slider4.leftPadding + slider4.visualPosition * (slider4.availableWidth - width)
                    y: slider4.topPadding + slider4.availableHeight / 2 - height / 2
                    color: slider4.pressed ? "#f0f0f0" : "#f6f6f6"
                    radius: 13
                    implicitHeight: 18
                    border.color: "#bdbebf"
                    implicitWidth: 18
                }
                snapMode: "SnapAlways"
                from: 0
                wheelEnabled: false
                to: 40
            }

            Slider {
                id: slider7
                x: 140
                y: 351
                width: 369
                height: 30
                background: Rectangle {
                    x: slider7.leftPadding
                    y: slider7.topPadding + slider7.availableHeight / 2 - height / 2
                    width: slider7.availableWidth
                    height: implicitHeight
                    color: "#bdbebf"
                    radius: 2
                    Rectangle {
                        width: slider7.visualPosition * parent.width
                        height: parent.height
                        color: Config.color_primary
                        radius: 2
                    }
                    implicitHeight: 4
                    implicitWidth: 200
                }
                stepSize: 5
                value: 10
                spacing: 0
                Text {
                    id: title7
                    x: 5
                    y: -20
                    text: "PEEP: "+slider7.value
                    font.pixelSize: 16
                    font.family: "Open Sans"
                }

                Text {
                    id: max7
                    x: 375
                    y: 7
                    text: slider7.to
                    font.pixelSize: 12
                }

                Text {
                    id: min7
                    x: -13
                    y: 7
                    text: slider7.from
                    font.pixelSize: 12
                }
                handle: Rectangle {
                    x: slider7.leftPadding + slider7.visualPosition * (slider7.availableWidth - width)
                    y: slider7.topPadding + slider7.availableHeight / 2 - height / 2
                    color: slider7.pressed ? "#f0f0f0" : "#f6f6f6"
                    radius: 13
                    implicitHeight: 18
                    border.color: "#bdbebf"
                    implicitWidth: 18
                }
                snapMode: "SnapAlways"
                from: 5
                wheelEnabled: false
                to: 20
            }


            Slider {
                id: slider8
                x: 143
                y: 413
                width: 369
                height: 30
                background: Rectangle {
                    x: slider8.leftPadding
                    y: slider8.topPadding + slider8.availableHeight / 2 - height / 2
                    width: slider8.availableWidth
                    height: implicitHeight
                    color: "#bdbebf"
                    radius: 2
                    Rectangle {
                        width: slider8.visualPosition * parent.width
                        height: parent.height
                        color: Config.color_primary
                        radius: 2
                    }
                    implicitHeight: 4
                    implicitWidth: 200
                }
                stepSize: 5
                value: 0.5
                spacing: 0
                Text {
                    id: title8
                    x: 5
                    y: -20
                    text: "FIO2: "+slider8.value + "%"
                    font.pixelSize: 16
                    font.family: "Open Sans"
                }

                Text {
                    id: max8
                    x: 375
                    y: 7
                    text: slider8.to
                    font.pixelSize: 12
                }

                Text {
                    id: min8
                    x: -13
                    y: 7
                    text: slider8.from
                    font.pixelSize: 12
                }
                handle: Rectangle {
                    x: slider8.leftPadding + slider8.visualPosition * (slider8.availableWidth - width)
                    y: slider8.topPadding + slider8.availableHeight / 2 - height / 2
                    color: slider8.pressed ? "#f0f0f0" : "#f6f6f6"
                    radius: 13
                    implicitHeight: 18
                    border.color: "#bdbebf"
                    implicitWidth: 18
                }
                snapMode: "SnapAlways"
                from: 21
                wheelEnabled: false
                to: 100
            }

            Rectangle {
                id: buttonSubmit
                x: 267
                y: 461
                width: 110
                height: 38
                color: Config.color_dark
                MouseArea{
                    anchors.fill: parent
                    Connections{
                        target: QmlBridge

                    }
                    onReleased: QmlBridge.sendPAC(2, slider2.value,slider3.value,slider4.value,slider7.value,slider8.value)
                }

                Text {
                    id: element
                    text: qsTr("SUBMIT")
                    font.bold: true
                    color: "white"
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                    anchors.fill: parent
                    font.pixelSize: 12
                }
            }



        }
        ButtonGroup { id: radioGroup }

    }
}


/*##^##
Designer {
    D{i:0;formeditorZoom:1.5}D{i:8;anchors_height:8;anchors_width:8;anchors_x:6;anchors_y:6}
D{i:9;anchors_height:8;anchors_width:8;anchors_x:6;anchors_y:6}D{i:52;anchors_height:8;anchors_width:8;anchors_x:6;anchors_y:6}
D{i:53;anchors_height:8;anchors_width:8;anchors_x:6;anchors_y:6}D{i:2;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/
