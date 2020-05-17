import QtQuick 2.0
import QtCharts 2.2
import QtQuick.Controls 2.0
import "./config.js"
as Config

Item {

    id: root

    property bool starttwo:false
    property int limit:50
    property int minY:0
    property int maxY:50
    property int xval:0
    property int xval2:-1
    signal reemitted(point p)

    Component.onCompleted: ChartManager.dataReady.connect(root.reemitted)
    onReemitted: {
        root.addpoint(p.y, xval, series1)
        xval++
        if (xval==limit){starttwo=true}
        // start second chart
        if(starttwo){
            root.addpoint2(p.y, xval2, series2)
            xval2++
        }
    }

    function addpoint(y,x, myseries) {
        if (x>limit){
            if(myseries.count!==0){
                myseries.remove(0)
            }else{
                xval=0
            }
        }else{
        myseries.append(x, y)
        }
    }

    function addpoint2(y,x, myseries) {
        if (x>limit){
            if(myseries.count!==0){
                myseries.remove(0)
            }else{
                xval2=0
            }
        }else{
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
                id: series2
                color: "blue"
                name: "LineSeries"

                axisY: ValueAxis {
                    id: axisY2
                    min: root.minY
                    max: root.maxY
                    minorTickCount: 1
                    tickCount: 3
                    visible:false
                }
                axisX: ValueAxis {
                    id: axisXs2
                    min: 0
                    max: root.limit
                    tickCount: 2
                    visible:false
                }

            }


            SplineSeries {
                id: series1
                color: "blue"
                name: "LineSeries"

                axisY: ValueAxis {
                    id: axisY
                    min: root.minY
                    max: root.maxY
                    minorTickCount: 1
                    tickCount: 3
                }
                axisX: ValueAxis {
                    id: axisXs
                    min: 0
                    max: root.limit
                    tickCount: 2
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
