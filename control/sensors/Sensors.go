package sensors

import "vent/pkg/adc"

// Pressure is a custom type struct to identify onboard
// pressure sensors
type Pressure struct {
	Name  string
	ID    int
	MMH2O float32
	AdcID uint8
}

// Flow is a custom type struct to identify onboard
// flow sensors
type Flow struct {
	Name  string
	ID    int
	Rate  float32
	AdcID uint8
}

// ReadPressure ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Pressure) ReadPressure() float32 {
	//Read ADC 1 (0-10 part), and return sllice of 8 bits
	AdcSlice, _ := adc.ReadADC(1)
	//used predefined id in struct as index to access data from slice
	VoltageSignal := AdcSlice[PS.ID]
	return PS.MMH2O
}

// ReadFlow ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Flow) ReadFlow() float32 {
	//Read ADC 1 (0-10 part), and return sllice of 8 bits
	AdcSlice, _ := adc.ReadADC(1)
	//used predefined id in struct as index to access data from slice
	VoltageSignal = AdcSlice[PS.ID]
	return PS.Rate
}
