import QtQuick 2.0

Item {
    property string iconColor
    property int iconSize: 32 // in pixcel
    property int level // 0-4 levels

    function getUnicode()
    {
        if (level == 0)
            return "\uf244" // fa-battery-empty
        else if (level == 1)
            return "\uf243" // fa-battery-quarter
        else if (level == 2)
            return "\uf242" // fa-battery-half
        else if (level == 3)
            return "\uf241" // fa-battery-three-quarters
        else if (level == 4)
            return "\uf240" // fa-battery-full
        else
            return "\uf244" // fa-battery-empty
    }

    Text {
        font.family: fontAwesome.name
        color: iconColor
        font.pixelSize: iconSize
        text: getUnicode()
    }

    FontLoader { id: fontAwesome; source: "./fa.ttf" }
}
