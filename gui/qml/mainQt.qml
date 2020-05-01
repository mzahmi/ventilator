import QtQuick 2.0
import QtQuick.Window 2.12
import "pages"
import "pages/MyConstants.js" as Constants
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0

Window {
    id: window
    visible: true
    width: 800
    height: 460
    title: qsTr("Hello World")
    color: Constants.mainbg
MainGo{
}
}

/*##^##
Designer {
    D{i:1;anchors_height:245}D{i:2;anchors_x:0;anchors_y:0}
}
##^##*/
