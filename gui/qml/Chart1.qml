import QtQuick 2.0
import QtCharts 2.3
import QtQuick.Controls 2.0
import "./config.js" as Config

Item {

    Rectangle {
        id: chartsarea
        color: "#ffffff"
        anchors.fill: parent
        property var xvalues : [0,0,0,0,0,0,0,0,0,0,0,0,0]
        property var yvalues : [0,0,0,0,0,0,0,0,0,0,0,0,0]
        property int i: 0


        ChartView {
            id: line
            anchors.fill: parent
            antialiasing: true
            legend.visible: false

            LineSeries {
                id: series
                color: "black"
                name: "LineSeries"
                axisY: ValueAxis {
                    id: axisY
                    min: -5
                    max: 40
                }
                axisX: ValueAxis {
                    id: axisXs
                    min: 0
                    max: 40
                }

                function populateSeries(){
                    var i;
                    for (i = 0; i < chartsarea.xvalues.length; i++){
                        console.log(count)
                        series.append(chartsarea.xvalues[count], chartsarea.yvalues[count])
                    }
                }

                function addpoint(y){
                    series.remove(0)
                    series.append(chartsarea.i, y)
                    chartsarea.i++
                    if (chartsarea.i >39){

                        chartsarea.i = 0
                    }
                }

                Component.onCompleted: populateSeries(chartsarea.yvalues)

                Connections{
                    target: QmlBridge
                    onSendToQml: series.addpoint(data)

                }
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
