package sensors

import (
	"github.com/mzahmi/ventilator/control/adc"
)

// Pressure is a custom type struct to identify onboard
// pressure sensors
type Pressure struct {
	Name  string
	ID    int
	AdcID uint8
	MMH2O float32
}

// Flow is a custom type struct to identify onboard
// flow sensors
type Flow struct {
	Name  string
	ID    int
	AdcID uint8
	Rate  float32
}

// ReadPressure ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Pressure) ReadPressure() float32 {
	//read raw data from source and convert to mmH2O pressure reading
	AdcSlice, _ := adc.ReadADC(PS.AdcID)
	VoltageSignal := AdcSlice[PS.ID]
	bar := (VoltageSignal - 0.0196) / 0.2802 //initail pressure sensor calibration
	PS.MMH2O = bar * 0.10197162129779        // conversion from bar to mmH2O
	return PS.MMH2O

}

// ReadFlow ... reads raw data from attached sensors to
// Memberane based on its address
func (FS *Flow) ReadFlow() float32 {
	//read raw data from source and convert to flow rate reading
	AdcSlice, _ := adc.ReadADC(FS.AdcID)
	VoltageSignal := AdcSlice[FS.ID]
	FS.Rate = VoltageSignal * 2
	return FS.Rate
}
