import QtQuick 2.9
import QtCharts 2.2
import QtQuick.Controls 1.4
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.5
import QtQuick.Controls.Material 2.12

ApplicationWindow {
    id: mainWindow
    width: 640
    height: 480
    title: qsTr("Simple ui")
    visible: true
    locale: locale

    property int controlsColor: Material.DeepPurple
    property int controlsAccent: Material.BlueGrey
    property real x: 0.0
    property int controlsElevation: 6
    property int paneElevation: 4

    signal reemitted(point p)
    Component.onCompleted: Manager.dataReady.connect(mainWindow.reemitted)
    onReemitted: {
        testXAxis.max = Math.max(testXAxis.max, p.x)
        testXAxis.min = Math.min(testXAxis.min, p.x)
        testYAxis.max = Math.max(testYAxis.max, p.y)
        testYAxis.min = Math.min(testYAxis.min, p.y)
        mainLine.append(p.x, p.y)
    }

    function drawPoint(xy) {
        mainLine.append(xy[0], xy[1])
        if (mainWindow.x >= testXAxis.max) {
            testXAxis.max = mainWindow.x;
        }
        if (py >= testYAxis.max) {
            testYAxis.max = py;
        }
        if (py <= testYAxis.min) {
            testYAxis.min = py;
        }
    }

    function clearLine() {
        mainLine.clear();
        mainLine.append(0, 0);
    }

    Pane {
        id: mainPanel
        anchors.fill: parent
        //Material.theme: Material.Dark

        RowLayout {
            id: mainRowLO
            anchors.fill: parent
            spacing: 15

            //Chart pane
            Pane {
                id: chartPane
                Material.elevation: paneElevation
                //Material.background: Material.Grey
                Layout.fillHeight: true
                Layout.fillWidth: true
                Layout.minimumHeight: 200
                Layout.minimumWidth: 400

                ChartView {
                    id: testChart
                    title: "Line"
                    anchors.fill: parent
                    antialiasing: true
                    LineSeries {
                        id: mainLine
                        name: "LineSeries"
                        axisX: ValueAxis {
                            id: testXAxis
                            min: 0.0
                            max: 2.0
                        }
                        axisY: ValueAxis {
                            id: testYAxis
                            min: 0.0
                            max: 2.0
                        }
                        XYPoint {
                            x: 0;y: 0
                        }
                    }
                }
            }

            Pane {
                id: controlsPane
                Material.elevation: paneElevation
                //Material.background: Material.Grey
                Layout.fillHeight: true
                Layout.fillWidth: true
                Layout.minimumHeight: 200
                Layout.minimumWidth: 200
                Layout.maximumWidth: 200

                ColumnLayout {
                    id: controlsColumnLO
                    anchors.fill: parent
                    spacing: 40

                    Label {
                        id: powerLabel
                        text: "Exponent"
                        Layout.topMargin: 40
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                    }

                    SpinBox {
                        id: powerNum
                        from: 0
                        value: 1
                        to: 5
                        stepSize: 1
                        width: 80
                        validator: DoubleValidator {
                            bottom: Math.min(powerNum.from, powerNum.to)
                            top: Math.max(powerNum.from, powerNum.to)
                        }
                        Material.foreground: controlsColor
                        Material.accent: controlsAccent
                        Material.elevation: controlsElevation
                        Layout.fillWidth: true
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                        editable: true
                        onValueChanged: function() {
                            Manager.power = value;
                        }
                    }

                    Label {
                        id: multiplierLabel
                        text: "Multiplier"
                        Layout.fillWidth: true
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                    }

                    Slider {
                        id: multiplierSlider
                        from: -50
                        value: 1
                        to: 50
                        Material.foreground: controlsColor
                        Material.accent: controlsAccent
                        Material.elevation: controlsElevation
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                        Layout.fillWidth: true
                        onValueChanged: function() {
                            Manager.multiplier = value;
                        }
                    }

                    Label {
                        id: multValueLabel
                        text: String(multiplierSlider.value)
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                    }

                    Label {
                        id: delayLable
                        text: "Delay[s]"
                        Layout.fillWidth: true
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                    }

                    Slider {
                        id: delaySlider
                        from: 0.05
                        value: 0.1
                        to: 1
                        stepSize: 0.01
                        Material.foreground: controlsColor
                        Material.accent: controlsAccent
                        Material.elevation: controlsElevation
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                        Layout.fillWidth: true
                        onValueChanged: function() {
                            Manager.delay = value;
                        }
                    }

                    Label {
                        id: incrementLable
                        text: "Increment"
                        Layout.fillWidth: true
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                    }

                    Slider {
                        id: incrementSlider
                        from: 1.0
                        value: 1.0
                        to: 5.0
                        stepSize: 0.01
                        Material.foreground: controlsColor
                        Material.accent: controlsAccent
                        Material.elevation: controlsElevation
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                        Layout.fillWidth: true
                        onValueChanged: function() {
                            Manager.xIncrement = value;
                        }
                    }

                    Item {
                        // spacer item
                        id: controlsSpacer
                        Layout.fillWidth: true
                        Layout.fillHeight: true
                        Pane {
                            anchors.fill: parent
                        } //; Material.background: Material.Light; Material.elevation: 4 } // to visualize the spacer
                    }

                    Button {
                        id: startPointBtn
                        text: "START"
                        Material.foreground: controlsColor
                        Material.accent: controlsAccent
                        Material.elevation: controlsElevation
                        Layout.fillWidth: true
                        Layout.leftMargin: 10
                        Layout.rightMargin: 10
                        onClicked: function() {
                            console.log(text);
                            if (text == "START") {
                                clearLine();
                                Manager.starter = true;
                                text = "STOP";
                            } else {
                                Manager.starter = false;
                                text = "START";
                            }
                        }
                    }
                }
            }
        }
    }
}