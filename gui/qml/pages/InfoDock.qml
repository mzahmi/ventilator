import QtQuick 2.0
import "basic"
import "MyConstants.js" as Constants

Item {
    Rectangle {
        id: rectangle
        color: Constants.mainbg
        anchors.fill: parent

        BreathingText{
            textSize: 36
            anchors.fill: parent

        }
    }

}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}D{i:2;anchors_width:616;anchors_x:79;anchors_y:58}
D{i:1;anchors_height:200;anchors_width:200}
}
##^##*/
