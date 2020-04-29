import QtQuick 2.0

Item {
    property string iconColor
    property int iconSize: 32 // in pixcel
    property int level // 0-4 levels

    FontLoader {
        id: fontAwesome
        source: "fa.ttf"
    }

    Text {
        font.family: fontAwesome.name
        color: iconColor
        font.pixelSize: iconSize
        text: "\uf083"
    }


}
