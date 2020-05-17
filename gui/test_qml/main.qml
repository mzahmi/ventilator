import QtQuick 2.0
import "./componentCreation.js"
as MyScript
import QtQuick.Window 2.0
import QtQuick.Controls 2.1
import QtQuick.Layouts 1.3


Window {
    id: window
    property int componentCount

    visible: true
    width: 800
    height: 480
    // visibility: "FullScreen"
    title: qsTr("Hello World")
    Component.onCompleted: {
        componentCount = MyScript.getSliders()
        submitButton.y = componentCount * 50
    }

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: 900
        Column {
            anchors.fill: parent
            Item {
                id: appWindow

            }

        }
        Button {
            id: submitButton
            text: "submit"
            onClicked: MyScript.iter()
        }
    }
}






/*##^##
Designer {
    D{i:2;anchors_height:100;anchors_width:100;anchors_x:222;anchors_y:192}
}
##^##*/
