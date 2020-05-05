import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.0
import "./config.js" as Config

Item {
    id: name
    height: 600


    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.fill: parent

        Flickable {
            id: flickable
            anchors.rightMargin: 0
            anchors.leftMargin: 0
            contentHeight: 600
            anchors.fill: parent

            Slider {
                id: slider1
                x: 136
                y: 85
                width: 369
                height: 30
                wheelEnabled: false
                spacing: 0
                value: 0.5
                from:10
                to:20
                stepSize: 5
                snapMode: "SnapAlways"

                background: Rectangle {
                    x: slider1.leftPadding
                    y: slider1.topPadding + slider1.availableHeight / 2 - height / 2
                    implicitWidth: 200
                    implicitHeight: 4
                    width: slider1.availableWidth
                    height: implicitHeight
                    radius: 2
                    color: "#bdbebf"

                    Rectangle {
                        width: slider1.visualPosition * parent.width
                        height: parent.height
                        color: Config.color_primary
                        radius: 2
                    }
                }

                handle: Rectangle {
                    x: slider1.leftPadding + slider1.visualPosition * (slider1.availableWidth - width)
                    y: slider1.topPadding + slider1.availableHeight / 2 - height / 2
                    implicitWidth: 18
                    implicitHeight: 18
                    radius: 13
                    color: slider1.pressed ? "#f0f0f0" : "#f6f6f6"
                    border.color: "#bdbebf"
                }

                Text {
                    id: title1
                    x: 5
                    y: -20
                    text: "X: "+slider1.value
                    font.family: "Open Sans"
                    font.pixelSize: 16
                }

                Text {
                    id: max1
                    x: 375
                    y: 7
                    text: slider1.to
                    font.pixelSize: 12
                }

                Text {
                    id: min1
                    x: -13
                    y: 7
                    text: slider1.from
                    font.pixelSize: 12
                }

            }

            Rectangle {
                id: buttonSubmit
                x: 266
                y: 550
                width: 110
                height: 38
                color: Config.color_dark

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

            Text {
                id: element1
                x: 38
                y: 14
                text: qsTr("Volume AC")
                font.pixelSize: 24
            }

            Slider {
                id: slider2
                x: 136
                y: 150
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
                value: 0.5
                stepSize: 5
                Text {
                    id: title2
                    x: 5
                    y: -20
                    text: "X: "+slider2.value
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
                from: 5
                wheelEnabled: false
                to: 50
            }

            Slider {
                id: slider3
                x: 137
                y: 216
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
                stepSize: 0.5
                value: 0
                spacing: 0
                Text {
                    id: title3
                    x: 5
                    y: -20
                    text: "X: "+slider3.value
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
                from: -5
                wheelEnabled: false
                to: 2
            }

            Slider {
                id: slider4
                x: 137
                y: 281
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
                value: 0.5
                stepSize: 5
                Text {
                    id: title4
                    x: 5
                    y: -20
                    text: "X: "+slider4.value
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
                from: 5
                wheelEnabled: false
                to: 50
            }

            Slider {
                id: slider5
                x: 137
                y: 346
                width: 369
                height: 30
                background: Rectangle {
                    x: slider5.leftPadding
                    y: slider5.topPadding + slider5.availableHeight / 2 - height / 2
                    width: slider5.availableWidth
                    height: implicitHeight
                    color: "#bdbebf"
                    radius: 2
                    Rectangle {
                        width: slider5.visualPosition * parent.width
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
                    id: title5
                    x: 5
                    y: -20
                    text: "X: "+slider5.value
                    font.pixelSize: 16
                    font.family: "Open Sans"
                }

                Text {
                    id: max5
                    x: 375
                    y: 7
                    text: slider5.to
                    font.pixelSize: 12
                }

                Text {
                    id: min5
                    x: -13
                    y: 7
                    text: slider5.from
                    font.pixelSize: 12
                }
                handle: Rectangle {
                    x: slider5.leftPadding + slider5.visualPosition * (slider5.availableWidth - width)
                    y: slider5.topPadding + slider5.availableHeight / 2 - height / 2
                    color: slider5.pressed ? "#f0f0f0" : "#f6f6f6"
                    radius: 13
                    implicitHeight: 18
                    border.color: "#bdbebf"
                    implicitWidth: 18
                }
                snapMode: "SnapAlways"
                from: 5
                wheelEnabled: false
                to: 50
            }

            Slider {
                id: slider6
                x: 136
                y: 412
                width: 369
                height: 30
                background: Rectangle {
                    x: slider6.leftPadding
                    y: slider6.topPadding + slider6.availableHeight / 2 - height / 2
                    width: slider6.availableWidth
                    height: implicitHeight
                    color: "#bdbebf"
                    radius: 2
                    Rectangle {
                        width: slider6.visualPosition * parent.width
                        height: parent.height
                        color: Config.color_primary
                        radius: 2
                    }
                    implicitHeight: 4
                    implicitWidth: 200
                }
                spacing: 0
                value: 0.5
                stepSize: 5
                Text {
                    id: title6
                    x: 5
                    y: -20
                    text: "X: "+slider6.value
                    font.pixelSize: 16
                    font.family: "Open Sans"
                }

                Text {
                    id: max6
                    x: 375
                    y: 7
                    text: slider6.to
                    font.pixelSize: 12
                }

                Text {
                    id: min6
                    x: -13
                    y: 7
                    text: slider6.from
                    font.pixelSize: 12
                }
                handle: Rectangle {
                    x: slider6.leftPadding + slider6.visualPosition * (slider6.availableWidth - width)
                    y: slider6.topPadding + slider6.availableHeight / 2 - height / 2
                    color: slider6.pressed ? "#f0f0f0" : "#f6f6f6"
                    radius: 13
                    implicitHeight: 18
                    border.color: "#bdbebf"
                    implicitWidth: 18
                }
                snapMode: "SnapAlways"
                from: 5
                wheelEnabled: false
                to: 50
            }

            Slider {
                id: slider7
                x: 137
                y: 477
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
                value: 0.5
                spacing: 0
                Text {
                    id: title7
                    x: 5
                    y: -20
                    text: "X: "+slider7.value
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
                to: 50
            }


        }

    }
}


/*##^##
Designer {
    D{i:0;formeditorZoom:1.5}D{i:2;anchors_height:300;anchors_width:300;anchors_x:88;anchors_y:128}
D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/
