import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.0
import QtQuick.Controls.Styles 1.4
import "./config.js" as Config

Item {
    id: name


    Rectangle {
        id: rectangle
        color: "#ffffff"
        anchors.fill: parent

        Slider {
            id: control
            x: 192
            y: 122
            wheelEnabled: false
            spacing: 0
            value: 0.5
            from:5
            to:50
            stepSize: 5
            snapMode: "SnapAlways"

            background: Rectangle {
                x: control.leftPadding
                y: control.topPadding + control.availableHeight / 2 - height / 2
                implicitWidth: 200
                implicitHeight: 4
                width: control.availableWidth
                height: implicitHeight
                radius: 2
                color: "#bdbebf"

                Rectangle {
                    width: control.visualPosition * parent.width
                    height: parent.height
                    color: Config.color_primary
                    radius: 2
                }
            }

            handle: Rectangle {
                x: control.leftPadding + control.visualPosition * (control.availableWidth - width)
                y: control.topPadding + control.availableHeight / 2 - height / 2
                implicitWidth: 18
                implicitHeight: 18
                radius: 13
                color: control.pressed ? "#f0f0f0" : "#f6f6f6"
                border.color: "#bdbebf"
            }
        }

        Slider {
            id: control1
            x: 192
            y: 202
            stepSize: 5
            wheelEnabled: false
            handle: Rectangle {
                x: control1.leftPadding + control1.visualPosition * (control1.availableWidth - width)
                y: control1.topPadding + control1.availableHeight / 2 - height / 2
                color: control1.pressed ? "#f0f0f0" : "#f6f6f6"
                radius: 13
                implicitWidth: 18
                border.color: "#bdbebf"
                implicitHeight: 18
            }
            value: 0.5
            spacing: 0
            background: Rectangle {
                x: control1.leftPadding
                y: control1.topPadding + control1.availableHeight / 2 - height / 2
                width: control1.availableWidth
                height: implicitHeight
                color: "#bdbebf"
                radius: 2
                Rectangle {
                    width: control1.visualPosition * parent.width
                    height: parent.height
                    color: Config.color_primary
                    radius: 2
                }
                implicitWidth: 200
                implicitHeight: 4
            }
            from: 5
            to: 50
            snapMode: "SnapAlways"
        }

        Slider {
            id: control2
            x: 192
            y: 285
            stepSize: 5
            wheelEnabled: false
            handle: Rectangle {
                x: control2.leftPadding + control2.visualPosition * (control2.availableWidth - width)
                y: control2.topPadding + control2.availableHeight / 2 - height / 2
                color: control2.pressed ? "#f0f0f0" : "#f6f6f6"
                radius: 13
                border.color: "#bdbebf"
                implicitWidth: 18
                implicitHeight: 18
            }
            value: 0.5
            spacing: 0
            background: Rectangle {
                x: control2.leftPadding
                y: control2.topPadding + control2.availableHeight / 2 - height / 2
                width: control2.availableWidth
                height: implicitHeight
                color: "#bdbebf"
                radius: 2
                Rectangle {
                    width: control2.visualPosition * parent.width
                    height: parent.height
                    color: Config.color_primary
                    radius: 2
                }
                implicitWidth: 200
                implicitHeight: 4
            }
            snapMode: "SnapAlways"
            to: 50
            from: 5
        }

        Text {
            id: element
            x: 197
            y: 102
            text: "PEEP: "+control.value
            font.pixelSize: 16
        }
    }
}


/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/
