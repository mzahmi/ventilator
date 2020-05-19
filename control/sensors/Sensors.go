//Package sensors ...
package sensors

import (
	// "fmt"
	"github.com/mzahmi/ventilator/control/adc"
)

//Inhalation pressure sensor constants
const Inhalation_Pressure_Sensor_Name = "Inhalation Pressure Sensor"
const Inhalation_Pressure_Sensor_ADC_ID = 1
const Inhalation_Pressure_Sensor_ADC_Chan = 1

//Inhalation flow sensor constants
const Inhalation_Flow_Sensor_Name = "Inhalation Flow Sensor"
const Inhalation_Flow_Sensor_ADC_ID = 1
const Inhalation_Flow_Sensor_ADC_Chan = 5

//Exhalaltion pressure sensor constants
const Exhalation_Pressure_Sensor_Name = "Exhalation Pressure Sensor"
const Exhalation_Pressure_Sensor_ADC_ID = 2
const Exhalation_Pressure_Sensor_ADC_Chan = 1

//Exhalaltion flow sensor constants
const Exhalation_Flow_Sensor_Name = "Exhalation Flow Sensor"
const Exhalation_Flow_Sensor_ADC_ID = 0
const Exhalation_Flow_Sensor_ADC_Chan = 0

// Pressure is a custom type struct to identify onboard
// pressure sensors
type Pressure struct {
	Name    string
	AdcID   uint8
	AdcChan uint
	Bar     float32
}

// Flow is a custom type struct to identify onboard
// flow sensors
type Flow struct {
	Name    string
	AdcID   uint8
	AdcChan uint
	Rate    float32
}

//SensorReading is custom type struct to store current sensor readings
type SensorsReading struct {
	PressureInput  float32
	PressureOutput float32
	FlowInput      float32
}

//PIns ... Inhalaltion pressure sensor
var PIns = Pressure{
	Name:    Inhalation_Pressure_Sensor_Name,
	AdcID:   Inhalation_Pressure_Sensor_ADC_ID,
	AdcChan: Inhalation_Pressure_Sensor_ADC_Chan,
	Bar:     0,
}

//FIns ... Inhalation flow sensor
var FIns = Flow{
	Name:    Inhalation_Flow_Sensor_Name,
	AdcID:   Inhalation_Pressure_Sensor_ADC_ID,
	AdcChan: Inhalation_Flow_Sensor_ADC_Chan,
	Rate:    0,
}

//PExp ... Exhalation pressure sensor
var PExp = Pressure{
	Name:    Exhalation_Pressure_Sensor_Name,
	AdcID:   Exhalation_Pressure_Sensor_ADC_ID,
	AdcChan: Exhalation_Pressure_Sensor_ADC_Chan,
	Bar:     0,
}

//FExp ... Exhalation flow sensor
var FExp = Flow{
	Name:    Exhalation_Flow_Sensor_Name,
	AdcID:   Exhalation_Flow_Sensor_ADC_ID,
	AdcChan: Exhalation_Flow_Sensor_ADC_Chan,
	Rate:    0,
}

// ReadPressureRaw ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Pressure) ReadPressureRaw() float32 {
	//read raw data from source and convert to mmH2O pressure reading
	AdcSlice, _ := adc.ReadADC(PS.AdcID)
	return AdcSlice[PS.AdcChan]

}

// ReadPressureBar ... reads data from attached sensors to
// Memberane based on its address using pressure
func (PS *Pressure) ReadPressure() float32 {
	PS.Bar = (PS.ReadPressureRaw()*1.37 + 0.11)
	return PS.Bar
}

//PressureToVolt ... converts pressure to volts
func PressureToVolt(pressure float32) float64 {
	volt := float64((pressure - .11) / 1.37)

	if volt < 0 {
		return 0
	} else {
		return volt
	}
}

// ReadFlow ... reads raw data from attached sensors to
// Memberane based on its address
func (FS *Flow) ReadFlow() float32 {
	//read raw data from source and convert to flow rate reading
	AdcSlice, _ := adc.ReadADC(FS.AdcID)
	VoltageSignal := AdcSlice[FS.AdcChan]
	return VoltageSignal
}

// // ReadAllSensors constantly reading input from the sensors and returns their readings
// // in this order: PIns, PExp, FIns, FExp readings
func ReadAllSensors() (InputPress, OutputPress, InputFlow float32) { //, InputFlow, OutputFlow float32) {
	// fmt.Println("read PIns")
	InputPress = PIns.ReadPressureRaw()
	// fmt.Println("read PExp")
	OutputPress = PExp.ReadPressureRaw()
	// fmt.Println("done Reading")
	InputFlow = FIns.ReadFlow()
	//OutputFlow = FExp.ReadFlow()
	return InputPress, OutputPress, InputFlow //OutputFlow
}

// func (PS *Pressure) CalibratePressure() {

// 	fmt.Println("Beginning Pressure Sensor Calibration")
// 	fmt.Println("Enter number of increments")
// 	increment := 0
// 	fmt.Scan(&increment)
// 	pressureCurve := make([]float32, increment)
// 	voltageCurve := make([]float32, increment)
// 	fmt.Println("Calibration loop starting. Set the pressure and record the values")
// 	response := "default"
// 	for ii := 0; ii <= increment; ii++ {
// 		for {
// 			fmt.Scan(&response)
// 			if response == "set" {
// 				voltageCurve[ii] = PS.ReadPressureRaw()
// 				fmt.Println("Enter the pressure in bar")
// 				fmt.Scan(&pressureCurve[ii])
// 				break
// 			}
// 		}
// 	}

// }
