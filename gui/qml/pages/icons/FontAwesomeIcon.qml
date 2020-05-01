import QtQuick 2.0

Item {
    property string iconColor
    property int iconSize: 32 // in pixcel
    property string iconStyle

    Text {
        font.family: fontAwesome.name
        color: iconColor
        font.pixelSize: iconSize
        text: iconStyle
        width: 24
        height: 24
    }

    FontLoader { id: fontAwesome; source: "fa.ttf" }
}
