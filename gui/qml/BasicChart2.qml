// volume chart
import QtQuick 2.0
import QtCharts 2.2
import QtQuick.Controls 2.0
import "./config.js"
as Config

Item {

    id: root

    property bool starttwo: false
    property double limit: 10.0
    property int minY: 0
    property double maxY: 1.5
    property double xval: 0
    property double xval2: 0
    signal reemitted(point p)

    Component.onCompleted: ChartManager2.dataReady.connect(root.reemitted)
    onReemitted: {
        root.addpoint(p.y, xval, series1)
        xval += 0.1
        if (xval > limit) {
            starttwo = true
        }
        // start second chart
        if (starttwo) {
            root.addpoint2(p.y, xval2, series2)
            xval2 += 0.1
        }
    }

    function addpoint(y, x, myseries) {
        if (x > limit) {
            if (myseries.count !== 0) {
                myseries.remove(0)
            } else {
                xval = 0
            }
        } else {
            myseries.append(x, y)
        }
    }

    function addpoint2(y, x, myseries) {
        if (x > limit) {
            if (myseries.count !== 0) {
                myseries.remove(0)
            } else {
                xval2 = 0
            }
        } else {
            myseries.append(x, y)
        }
    }

    Rectangle {
        id: chartsarea
        color: "#ffffff"
        anchors.fill: parent

        ChartView {
            id: chartview
            anchors.fill: parent
            antialiasing: true
            legend.visible: false

            SplineSeries {
                id: series1
                color: "#5677fc"
                name: "LineSeries"
                width: 1.5

                axisY: ValueAxis {
                    id: axisY
                    min: root.minY
                    max: root.maxY
                    minorTickCount: 1
                    tickCount: 3
                    titleText: "L"

                }
                axisX: ValueAxis {
                    id: axisXs
                    min: 0
                    max: root.limit
                    tickCount: 2
                }

            }

            SplineSeries {
                id: series2
                color: "#5677fc"
                name: "LineSeries"
                width: 1.5

                axisY: ValueAxis {
                    id: axisY2
                    min: root.minY
                    max: root.maxY
                    minorTickCount: 1
                    tickCount: 3
                    visible: false
                }
                axisX: ValueAxis {
                    id: axisXs2
                    min: 0
                    max: root.limit
                    tickCount: 2
                    visible: false
                }

            }
        }
    }

}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}
}
##^##*/
