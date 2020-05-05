import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.0
import "./config.js" as Config
import "./ViewModeSelector.js" as VMS

Item {
    id: name
    width: 650
    height: 460

    StackView {
        id: view
        anchors.fill: parent
        initialItem: selectmode

        Component{
            id: selectmode
            Item{
                id: element1
                anchors.fill: parent
                Text {
                    id: title
                    y: 78
                    text: qsTr("Select Mode")
                    horizontalAlignment: Text.AlignHCenter
                    anchors.left: parent.left
                    anchors.right: parent.right
                    font.pointSize: 32
                }



                RowLayout {
                    y: 180
                    anchors.rightMargin: 20
                    anchors.leftMargin: 20
                    anchors.right: parent.right
                    anchors.left: parent.left

                    Button{
                        text: "Volume AC"
                        onClicked: view.push(selectbreathe)
                    }

                    Button {
                        id: button2
                        text: "Pressure AC"
                        onClicked: view.push(selectbreathe)
                    }
                    Button {
                        id: button3
                        text: "PSV"
                        onClicked: view.push(selectbreathe)
                    }
                    Button {
                        id: button4
                        text: "P-SIMV"
                        onClicked: view.push(selectbreathe)
                    }
                    Button {
                        id: button5
                        text: "S-SIMV"
                        onClicked: view.push(selectbreathe)
                    }
                }
            }
        }

        Component{
            id: selectbreathe
            Item{
                id: element
                anchors.fill: parent
                Text {
                    id: title
                    y: 78
                    text: qsTr("Select Breathe Type")
                    horizontalAlignment: Text.AlignHCenter
                    anchors.left: parent.left
                    anchors.right: parent.right
                    font.pointSize: 32
                }
                Button{
                    id: button1
                    text: "<"
                    onClicked: view.pop()
                }

                RowLayout {
                    x: 223
                    y: 233

                    Button {
                        id: button
                        text: qsTr("Assist")

                    }

                    Button {
                        id: button6
                        text: qsTr("Control")
                        onClicked: view.push(selecttrigger)
                    }
                }
            }
        }

        Component{
            id: selecttrigger
            Flickable{
                contentHeight:800
                anchors.fill: parent
            S_m2b1t0{

            }
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
