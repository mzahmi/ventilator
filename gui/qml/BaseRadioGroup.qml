import QtQuick 2.0
import "./config.js"
as Config
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0

Item {
    id: root
    property string name: "IE"
    property int value: 2
    x: 50

    Rectangle {
        id: rectangle
        width: 490
        height: 60
        color: "#ffffff"

        ButtonGroup {
            id: radioGroup
        }

        ColumnLayout {
            y: 0
            anchors.left: parent.left
            anchors.leftMargin: 0
            anchors.right: parent.right

            Text {
                id: ieratiotitle
                text: root.name
                font.pixelSize: 16
                font.family: "Open Sans"
            }

            Row {
                Layout.fillWidth: true
                Layout.preferredHeight: 40

                spacing: 100

                RadioButton {
                    id: control
                    text: qsTr("1:1")
                    ButtonGroup.group: radioGroup
                    onClicked: {
                        root.value = 1
                    }

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
                    ButtonGroup.group: radioGroup
                    checked: true
                    onClicked: {
                        root.value = 2
                    }

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
                    ButtonGroup.group: radioGroup
                    onClicked: {
                        root.value = 3
                    }
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
                    onClicked: {
                        root.value = 4
                    }
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
        }
    }
}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:2;anchors_x:0}D{i:1;anchors_width:490;anchors_x:75}
}
##^##*/

