package sensors

// Pressure is a custom type struct to identify onboard
// pressure sensors
type Pressure struct {
	ID       string
	Address  string
	RawValue float64
	MMH2O    float64
}

// Flow is a custom type struct to identify onboard
// flow sensors
type Flow struct {
	ID       string
	Address  string
	RawValue float64
	Rate     float64
}

// ReadPressure ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Pressure) ReadPressure() float64 {
	//read raw data from source and convert to mmH2O pressure reading

	return PS.MMH2O

}

// ReadFlow ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Flow) ReadFlow() float64 {
	//read raw data from source and convert to flow rate reading
	return PS.Rate
}
