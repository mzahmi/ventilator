import QtQuick 2.0
import QtCharts 2.2
import QtQuick.Controls 2.0
import "./config.js"
as Config

Item {

    id: mainItem
    signal reemitted(point p)
    // connects to reemitted
    Component.onCompleted: ChartManager.dataReady.connect(mainItem.reemitted)
    onReemitted: {
        series1.addpoint(p.y)
    }

    Rectangle {
        id: chartsarea
        color: "#ffffff"
        anchors.fill: parent
        property var xvalues: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        property var yvalues: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
        property int i: 0
        property int j: 0
        property bool seriesswitch: false


        ChartView {
            id: chartview
            anchors.fill: parent
            antialiasing: true
            legend.visible: false


            LineSeries {
                id: series1
                color: "black"
                name: "LineSeries"

                axisY: ValueAxis {
                    id: axisY
                    min: -5
                    max: 50
                    minorTickCount: 1
                    tickCount: 3

                }
                axisX: ValueAxis {
                    id: axisXs
                    min: 0
                    max: 40
                }

                function populateSeries(myseries) {
                    var i;
                    for (i = 0; i < chartsarea.xvalues.length; i++) {
                        myseries.append(chartsarea.xvalues[count], chartsarea.yvalues[count])
                    }
                }

                function addpoint(y) {
                    series1.remove(0)
                    series1.append(chartsarea.i, y)
                    chartsarea.i++
                    if (chartsarea.i > 40) {
                        axisXs.min = axisXs.min + 1
                        axisXs.max = axisXs.max + 1
                    }
                }

                Component.onCompleted: {
                    populateSeries(series1)
                }

                // Connections{
                //     target: QmlBridge
                //     onSendToQml: series1.addpoint(data)

                // }
            }
        }
        //        Button {
        //            id: button
        //            x: 61
        //            y: 30
        //            text: qsTr("Button")
        //            property int counter: 5
        //            onClicked: {
        //                series.addpoint(counter,counter)
        //                counter++
        //            }
        //        }
    }

}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}
}
##^##*/
