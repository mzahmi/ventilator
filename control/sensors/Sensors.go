package sensors

import (
	"github.com/mzahmi/ventilator/control/adc"
)

//Inhalation pressure sensor constants
const Inhalation_Pressure_Sensor_Name = "Inhalation Pressure Sensor"
const Inhalation_Pressure_Sensor_ADC_ID = 1
const Inhalation_Pressure_Sensor_ADC_Chan = 0

//Inhalation flow sensor constants
const Inhalation_Flow_Sensor_Name = "Inhalation Flow Sensor"
const Inhalation_Flow_Sensor_ADC_ID = 1
const Inhalation_Flow_Sensor_ADC_Chan = 1

//Exhalaltion pressure sensor constants
const Exhalation_Pressure_Sensor_Name = "Exhalation Pressure Sensor"
const Exhalation_Pressure_Sensor_ADC_ID = 1
const Exhalation_Pressure_Sensor_ADC_Chan = 2

//Exhalaltion flow sensor constants
const Exhalation_Flow_Sensor_Name = "Exhalation Flow Sensor"
const Exhalation_Flow_Sensor_ADC_ID = 1
const Exhalation_Flow_Sensor_ADC_Chan = 3

// Pressure is a custom type struct to identify onboard
// pressure sensors
type Pressure struct {
	Name    string
	AdcID   uint8
	AdcChan uint
	MMH2O   float32
}

// Flow is a custom type struct to identify onboard
// flow sensors
type Flow struct {
	Name    string
	AdcID   uint8
	AdcChan uint
	Rate    float32
}

//PIns ... Inhalaltion pressure sensor
var PIns = Pressure{
	Name:    Inhalation_Pressure_Sensor_Name,
	AdcID:   Inhalation_Pressure_Sensor_ADC_ID,
	AdcChan: Inhalation_Pressure_Sensor_ADC_Chan,
	MMH2O:   0,
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
	MMH2O:   0,
}

//FExp ... Exhalation flow sensor
var FExp = Flow{
	Name:    Exhalation_Flow_Sensor_Name,
	AdcID:   Exhalation_Flow_Sensor_ADC_ID,
	AdcChan: Exhalation_Flow_Sensor_ADC_Chan,
	Rate:    0,
}

// ReadPressure ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Pressure) ReadPressure() float32 {
	//read raw data from source and convert to mmH2O pressure reading
	AdcSlice, _ := adc.ReadADC(PS.AdcID)
	VoltageSignal := AdcSlice[PS.AdcChan]
	bar := (VoltageSignal - 0.0196) / 0.2802 //initail pressure sensor calibration
	PS.MMH2O = bar * 0.10197162129779        // conversion from bar to mmH2O
	return PS.MMH2O

}

// ReadFlow ... reads raw data from attached sensors to
// Memberane based on its address
func (FS *Flow) ReadFlow() float32 {
	//read raw data from source and convert to flow rate reading
	AdcSlice, _ := adc.ReadADC(FS.AdcID)
	VoltageSignal := AdcSlice[FS.AdcChan]
	FS.Rate = VoltageSignal * 2
	return FS.Rate
}
