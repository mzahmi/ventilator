import QtQuick 2.0
import QtQuick.Controls 2.0
import "../qml/config.js" as Config

Item {
    id: root
    property string name: "PIP"
    property int initialVal: 25
    property int minVal: 15
    property int maxVal: 40
    property int stepSize: 5
    property alias value: slider.value
    width: children[0].width
    height: children[0].height
    property bool active:true
    signal activated()
    Component.onCompleted:{
        ModeSelect.modeSelected.connect(root.activated)
    }
    onActivated:{
        root.active = ModeSelect.status==="running"?false:true
    }

    Rectangle {
        id: rectangle
        width: slider.width+120
        height: slider.height+50
        color: "#ffffff"

        Slider {
            id: slider
            value: root.initialVal
            stepSize: root.stepSize
            from: root.minVal
            to: root.maxVal

            x:50
            y:30
            spacing: 0
            width: 369
            height: 30
            snapMode: "SnapAlways"
            wheelEnabled: false

            // indicator
            background: Rectangle {
                x: slider.leftPadding
                y: (slider.topPadding + slider.availableHeight / 2 - height / 2)
                width: slider.availableWidth
                height: implicitHeight
                color: "#bdbebf"
                radius: 2
                implicitHeight: 4
                implicitWidth: 200


                Rectangle {
                    width: slider.visualPosition * parent.width
                    height: parent.height
                    color: Config.color_primary
                    radius: 2
                }
            }

            // color
            handle: Rectangle {
                x: slider.leftPadding + slider.visualPosition * (slider.availableWidth - width)
                y: slider.topPadding + slider.availableHeight / 2 - height / 2
                color: slider.pressed ? "#f0f0f0" : "#f6f6f6"
                radius: 13
                implicitHeight: 18
                border.color: "#bdbebf"
                implicitWidth: 18
            }


            // title appears ontop of slider
            Text {
                x: 5
                y: -20
                color:root.active?"black":"grey"
                text: root.name+ ": "+slider.value
                font.pixelSize: 16
                font.family: "Open Sans"
            }

            // max value
            Text {
                x: 375
                y: 7
                color:root.active?"black":"grey"
                text: slider.to
                font.pixelSize: 12
            }

            // minimum value
            Text {
                id: min3
                x: -25
                y: 7
                width: 26
                height: 17
                color:root.active?"black":"grey"
                text: slider.from
                horizontalAlignment: Text.AlignRight
                font.pixelSize: 12
            }

            // minus button
            MouseArea {
                id: ma3
                x: -59
                y: 0
                width: 40
                height: 30
                onClicked: {
                    slider.value=slider.value-slider.stepSize
                }
                Text {
                    id: minus2
                    color: root.active?Config.color_dark:"grey"
                    text: "\uf068"
                    
                    anchors.fill: parent
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                    font.pixelSize: 18
                }
            }

            // plus button
            MouseArea {
                id: ma4
                x: 395
                y: 0
                width: 40
                height: 30
                onClicked: {
                    slider.value=slider.value+slider.stepSize
                }
                Text {
                    id: plus2
                    color: root.active?Config.color_dark:"grey"
                    text: "\uf067"
                    
                    anchors.fill: parent
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                    font.pixelSize: 18
                }
            }
        }
    }
}

/*##^##
Designer {
    D{i:0;formeditorZoom:2}D{i:2;anchors_height:30;anchors_width:369}
}
##^##*/
