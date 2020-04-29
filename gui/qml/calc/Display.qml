// Copyright (C) 2015 The Qt Company Ltd.
import QtQuick 2.0
import QtQuick.Window 2.0    
Item {
    id: display
    property real fontSize: Math.floor(Screen.pixelDensity * 10.0)
    property string fontColor: "#000000"
    property bool enteringDigits: false
    property int maxDigits: (width / fontSize) + 1
    property string displayedOperand
    property string errorString: qsTr("ERROR")
    property bool isError: displayedOperand === errorString

    function displayOperator(operator)
    {
        listView.model.append({ "operator": operator, "operand": "" })
        enteringDigits = true
        listView.positionViewAtEnd()
        //console.log("display",operator);
    }

    function newLine(operator, operand)
    {
        displayedOperand = displayNumber(operand)
        listView.model.append({ "operator": operator, "operand": displayedOperand })
        enteringDigits = false
        listView.positionViewAtEnd()
        //console.log("newLine",operator);
    }

    function appendDigit(digit)
    {
        if (!enteringDigits)
            listView.model.append({ "operator": "", "operand": "" })
        var i = listView.model.count - 1;
        listView.model.get(i).operand = listView.model.get(i).operand + digit;
        enteringDigits = true
        listView.positionViewAtEnd()
        //console.log("num is ", digit);
    }

    function setDigit(digit)
    {
        var i = listView.model.count - 1;
        listView.model.get(i).operand = digit;
        listView.positionViewAtEnd()
        //console.log("setDigit",digit);
    }

    function clear()
    {
        displayedOperand = ""
        if (enteringDigits) {
            var i = listView.model.count - 1
            if (i >= 0)
                listView.model.remove(i)
            enteringDigits = false
        }
        //console.log("clearing");
    }

    // Returns a string representation of a number that fits in
    // display.maxDigits characters, trying to keep as much precision
    // as possible. If the number cannot be displayed, returns an
    // error string.
    function displayNumber(num) {
        if (typeof(num) != "number")
            return errorString;

        var intNum = parseInt(num);
        var intLen = intNum.toString().length;

        // Do not count the minus sign as a digit
        var maxLen = num < 0 ? maxDigits + 1 : maxDigits;

        if (num.toString().length <= maxLen) {
            if (isFinite(num))
                return num.toString();
            return errorString;
        }

        // Integer part of the number is too long - try
        // an exponential notation
        if (intNum == num || intLen > maxLen - 3) {
            var expVal = num.toExponential(maxDigits - 6).toString();
            if (expVal.length <= maxLen)
                return expVal;
        }

        // Try a float presentation with fixed number of digits
        var floatStr = parseFloat(num).toFixed(maxDigits - intLen - 1).toString();
        if (floatStr.length <= maxLen)
            return floatStr;

        return errorString;
    }

    Item {
        id: theItem
        width: parent.width
        height: parent.height

        Rectangle {
            id: rect
            x: 0
            color: "#eceeea"
            height: parent.height
            width: display.width
        }
        /*Image {
            anchors.right: rect.left
            source: "images/paper-edge-left.png"
            height: parent.height
            fillMode: Image.TileVertically
        }
        Image {
            anchors.left: rect.right
            source: "images/paper-edge-right.png"
            height: parent.height
            fillMode: Image.TileVertically
        }

        Image {
            id: grip
            source: "images/paper-grip.png"
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 20
        }*/

        ListView {
            id: listView
            width: display.width
            height: display.height
            delegate: Item {
                height: display.fontSize * 1.1
                width: parent.width
                Text {
                    id: operator
                    font.pixelSize: display.fontSize
                    color: "#6da43d"
                    text: model.operator
                }
                Text {
                    id: operand
                    y:5
                    font.pixelSize: display.fontSize
                    color: display.fontColor
                    anchors.right: parent.right
                    anchors.rightMargin: 22
                    text: model.operand
                }
            }
            model: ListModel { }
        }

    }

}