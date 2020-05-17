import QtQuick 2.0
import QtQuick.Layouts 1.0
import QtQuick.Controls 2.1

Item {
    anchors.fill: parent
    Rectangle {
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 0
        anchors.fill: parent

        GridLayout {
            x: 54
            y: 107
            columnSpacing: 40
            rowSpacing: 15
            rows: 5
            columns: 2

            Text {
                id: element
                text: qsTr("Name: ")
                font.pixelSize: 24
            }

            Text {
                id: element5
                text: Patient.name
                font.pixelSize: 24
            }


            Text {
                id: element1
                text: qsTr("Age:")
                font.pixelSize: 24
            }

            SpinBox {
                id: spinBox
                value: 30
            }

            Text {
                id: element2
                text: qsTr("Gender: ")
                font.pixelSize: 24
            }

            Switch {
                id: element6
                text: qsTr("Switch")
            }

            Text {
                id: element3
                text: qsTr("Height:")
                font.pixelSize: 24
            }

            SpinBox {
                id: spinBox2
                from: 0
                to: 300
                value: 160
            }

            Text {
                id: element4
                text: qsTr("Weight:")
                font.pixelSize: 24
            }

            SpinBox {
                id: spinBox3
                from: 0
                to: 300
                value: 70
            }





        }

        Text {
            id: element10
            x: 54
            y: 22
            width: 289
            height: 38
            text: qsTr("Patient Information")
            font.pixelSize: 28
        }


    }

}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:2;anchors_x:100;anchors_y:184}
}
##^##*/
