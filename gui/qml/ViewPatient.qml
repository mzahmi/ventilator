import QtQuick 2.0
import QtQuick.Layouts 1.0
import QtQuick.Controls 2.1
import QtQuick.Controls.Styles 1.4
import "./config.js"
as Config

Item {
    anchors.fill: parent
    Rectangle {
        id: rectangle
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 0
        anchors.fill: parent

        Text {
            id: element
            y: 15
            color: "#646464"
            text: qsTr("Patient Information")
            anchors.rightMargin: -11
            anchors.leftMargin: 51
            anchors.left: parent.left
            anchors.right: parent.right
            font.pixelSize: 28
        }

        Text {
            id: element2
            x: 51
            y: 52
            text: qsTr("Tell us about yourself")
            font.bold: true
            font.pixelSize: 16
        }

        Text {
            id: element3
            x: 51
            y: 129
            color: "#6c6c6c"
            text: qsTr("What is your gender?")
            font.pixelSize: 15
        }

        Rectangle {
            id: male
            property bool active: false
            x: 51
            y: 163
            width: 106
            height: 41
            color: "#ffffff"
            radius: 8
            border.color: active ? Config.color_primary : "grey"


            Text {
                y: 15
                color: male.active ? Config.color_primary : "grey"
                text: qsTr("Male")
                font.bold: true
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                anchors.right: parent.right
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }

            MouseArea {
                id: mouseAreaMale
                anchors.fill: parent
                onClicked: {
                    male.active = true
                    female.active = false
                }
            }
        }

        Rectangle {
            id: female
            property bool active: false
            x: 184
            y: 163
            width: 106
            height: 41
            color: "#ffffff"
            radius: 8
            border.color: active ? Config.color_primary : "grey"

            Text {
                y: 15
                color: female.active ? Config.color_primary : "grey"
                text: qsTr("Female")
                font.bold: true
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                anchors.right: parent.right
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }

            MouseArea {
                id: mouseAreaFemale
                anchors.fill: parent
                onClicked: {
                    female.active = true
                    male.active = false
                }
            }
        }

        Text {
            id: element4
            x: 352
            y: 129
            color: "#6c6c6c"
            text: qsTr("How old are you?")
            font.pixelSize: 15
        }


        Slider {
            id: sliderage
            x: 352
            y: 170
            width: 239
            height: 27
            value: 33
            from: 5
            to: 90

            background: Rectangle {
                x: sliderage.leftPadding
                y: sliderage.topPadding + sliderage.availableHeight / 2 - height / 2
                implicitWidth: 200
                implicitHeight: 4
                width: sliderage.availableWidth
                height: implicitHeight
                radius: 2
                color: "#bdbebf"

                Rectangle {
                    width: sliderage.visualPosition * parent.width
                    height: parent.height
                    color: Config.color_primary
                    radius: 2
                }
            }

            handle: Rectangle {
                x: sliderage.leftPadding + sliderage.visualPosition * (sliderage.availableWidth - width)
                y: sliderage.topPadding + sliderage.availableHeight / 2 - height / 2
                implicitWidth: 8
                implicitHeight: 15
                radius: 2
                color: sliderage.pressed ? "#f0f0f0" : "#f6f6f6"
                border.color: "#bdbebf"
            }
        }

        Text {
            id: agetext
            x: 501
            y: 129
            color: "grey"
            text: Math.floor(sliderage.value) + " years old"
            font.pixelSize: 14
        }

        Text {
            id: element5
            x: 51
            y: 272
            color: "#6c6c6c"
            text: qsTr("What is your height?")
            font.pixelSize: 15
        }

        Slider {
            id: sliderHeight
            x: 51
            y: 307
            width: 239
            height: 27
            to: 250
            from: 80
            value: 160
            background: Rectangle {
                x: sliderHeight.leftPadding
                y: sliderHeight.topPadding + sliderHeight.availableHeight / 2 - height / 2
                width: sliderHeight.availableWidth
                height: implicitHeight
                color: "#bdbebf"
                radius: 2
                implicitWidth: 200
                implicitHeight: 4
                Rectangle {
                    width: sliderHeight.visualPosition * parent.width
                    height: parent.height
                    color: Config.color_primary
                    radius: 2
                }
            }
            handle: Rectangle {
                x: sliderHeight.leftPadding + sliderHeight.visualPosition * (sliderHeight.availableWidth - width)
                y: sliderHeight.topPadding + sliderHeight.availableHeight / 2 - height / 2
                color: sliderHeight.pressed ? "#f0f0f0" : "#f6f6f6"
                radius: 2
                implicitWidth: 8
                implicitHeight: 15
                border.color: "#bdbebf"
            }

        }

        Text {
            id: agetext1
            x: 235
            y: 274
            color: "#808080"
            text: Math.floor(sliderHeight.value) + "cm"
            font.pixelSize: 14
        }

        Text {
            id: element6
            x: 352
            y: 272
            color: "#6c6c6c"
            text: qsTr("What is your weight?")
            font.pixelSize: 15
        }

        Slider {
            id: sliderWeight
            x: 352
            y: 307
            width: 239
            height: 27
            to: 200
            background: Rectangle {
                x: sliderWeight.leftPadding
                y: sliderWeight.topPadding + sliderWeight.availableHeight / 2 - height / 2
                width: sliderWeight.availableWidth
                height: implicitHeight
                color: "#bdbebf"
                radius: 2
                implicitWidth: 200
                implicitHeight: 4
                Rectangle {
                    width: sliderWeight.visualPosition * parent.width
                    height: parent.height
                    color: Config.color_primary
                    radius: 2
                }
            }
            handle: Rectangle {
                x: sliderWeight.leftPadding + sliderWeight.visualPosition * (sliderWeight.availableWidth - width)
                y: sliderWeight.topPadding + sliderWeight.availableHeight / 2 - height / 2
                color: sliderWeight.pressed ? "#f0f0f0" : "#f6f6f6"
                radius: 2
                implicitWidth: 8
                border.color: "#bdbebf"
                implicitHeight: 15
            }
            from: 40
            value: 80
        }

        Text {
            id: agetext2
            x: 536
            y: 274
            width: 44
            height: 20
            color: "#808080"
            text: Math.floor(sliderWeight.value) + "kg"
            verticalAlignment: Text.AlignBottom
            horizontalAlignment: Text.AlignRight
            font.pixelSize: 14
        }


        Rectangle {
            id: continueButton
            property bool active: false
            x: 267
            y: 390
            width: 106
            height: 41
            color: Config.color_primary
            radius: 8

            Text {
                y: 15
                color: "white"
                text: qsTr("Continue")
                font.bold: true
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                anchors.right: parent.right
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 12
            }

            MouseArea {
                id: mouseAreacontinue
                anchors.fill: parent
                onClicked: {

                }
            }
        }


    }

}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:2;anchors_x:32}D{i:6;anchors_height:10;anchors_x:8;anchors_y:16}
D{i:9;anchors_x:45}D{i:10;anchors_height:100;anchors_width:100}
}
##^##*/
