package sensors

// Pressure is a custom type struct to identify onboard
// pressure sensors
type Pressure struct {
	ID      string
	Address string
	MMH2O   float32
}

// Flow is a custom type struct to identify onboard
// flow sensors
type Flow struct {
	ID      string
	Address string
	Rate    float32
}

// ReadPressure ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Pressure) ReadPressure() float32 {
	//read raw data from source and convert to mmH2O pressure reading
	return PS.MMH2O

}

// ReadFlow ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Flow) ReadFlow() float32 {
	//read raw data from source and convert to flow rate reading
	return PS.Rate
}
