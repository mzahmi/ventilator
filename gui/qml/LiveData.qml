import QtQuick 2.0
import "./src/variables/fontawesome.js"
as FontAwesome
import "./src/lists"
import "./config.js"
as Config
import QtQuick.Layouts 1.0
import "."
import QtQuick.Controls 2.1


Item {
    id: monitorrow
    property int pip: 25
    property int vt: 10
    property int rate: 15
    property int peep: 4
    property int fio2: 25
    property string mode: "PAC"
    RowLayout {
        x: 0
        y: 178
        anchors.rightMargin: 0
        anchors.leftMargin: 0
        spacing: 0
        anchors.left: parent.left
        anchors.right: parent.right
        ColumnLayout {
            width: Config.sidebar_height / 2
            id: column5
            Text {
                id: element14
                color: "#ffffff"
                text: monitorrow.fio2
                font.pixelSize: Config.sidebar_livetext_size1
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element15
                color: "#ffffff"
                text: qsTr("FIO2")
                font.pixelSize: Config.sidebar_livetext_size2
                styleColor: "#ffffff"
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element16
                color: "#f9f9f9"
                text: qsTr("%")
                font.pixelSize: Config.sidebar_livetext_size3
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }
        }
        
        ColumnLayout {
            width: Config.sidebar_height / 2
            id: column6
            Text {
                id: element20
                color: "#ffffff"
                text: monitorrow.mode
                font.pixelSize: Config.sidebar_livetext_size1
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element21
                color: "#ffffff"
                text: qsTr("Mode")
                font.pixelSize: Config.sidebar_livetext_size2
                styleColor: "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element22
                color: "#f9f9f9"
                font.pixelSize: Config.sidebar_livetext_size3
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
        }
    }
    
    RowLayout {
        x: 0
        y: 90
        anchors.rightMargin: 0
        anchors.leftMargin: 0
        spacing: 0
        anchors.left: parent.left
        anchors.right: parent.right
        ColumnLayout {
            id: column3
            Text {
                id: element8
                color: "#ffffff"
                text: monitorrow.rate
                font.pixelSize: Config.sidebar_livetext_size1
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element9
                color: "#ffffff"
                text: qsTr("Rate")
                font.pixelSize: Config.sidebar_livetext_size2
                styleColor: "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element10
                color: "#f9f9f9"
                text: qsTr("BPM")
                font.pixelSize: Config.sidebar_livetext_size3
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
        }
        
        ColumnLayout {
            id: column4
            width: Config.sidebar_height / 2
            Text {
                id: element11
                color: "#ffffff"
                text: monitorrow.peep
                font.pixelSize: Config.sidebar_livetext_size1
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element12
                color: "#ffffff"
                text: qsTr("PEEP")
                font.pixelSize: Config.sidebar_livetext_size2
                styleColor: "#ffffff"
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element13
                color: "#f9f9f9"
                text: qsTr("cmH2O")
                font.pixelSize: Config.sidebar_livetext_size3
                Layout.fillWidth: true
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
            }
        }
    }
    
    Row {
        x: 0
        
        y: 3
        anchors.rightMargin: 0
        anchors.leftMargin: 0
        spacing: 0
        anchors.left: parent.left
        anchors.right: parent.right
        
        
        
        Column {
            id: column1
            width: Config.sidebar_width / 2
            Layout.fillHeight: false
            Layout.fillWidth: false
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            
            Text {
                id: piptext
                color: "#ffffff"
                text: monitorrow.pip
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: Config.sidebar_livetext_size1
            }
            
            Text {
                id: element2
                color: "#ffffff"
                text: qsTr("PIP")
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                styleColor: "#ffffff"
                font.pixelSize: Config.sidebar_livetext_size2
            }
            
            Text {
                id: element3
                color: "#f9f9f9"
                text: qsTr("cmH2O")
                anchors.right: parent.right
                anchors.left: parent.left
                anchors.leftMargin: 0
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: Config.sidebar_livetext_size3
            }
        }
        
        Column {
            id: column2
            width: Config.sidebar_width / 2
            Layout.fillWidth: false
            Text {
                id: element4
                color: "#ffffff"
                text: monitorrow.vt
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                font.pixelSize: Config.sidebar_livetext_size1
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element5
                color: "#ffffff"
                text: qsTr("Vt")
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                font.pixelSize: Config.sidebar_livetext_size2
                styleColor: "#ffffff"
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
            
            Text {
                id: element7
                color: "#f9f9f9"
                text: qsTr("ml")
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                font.pixelSize: Config.sidebar_livetext_size3
                horizontalAlignment: Text.AlignHCenter
                Layout.fillWidth: true
                verticalAlignment: Text.AlignVCenter
            }
        }
    }
    
    
}

