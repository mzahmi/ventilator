package sensors

// Pressure is a custom type struct to identify onboard
// pressure sensors
type Pressure struct {
	ID       int
	Address  int
	RawValue float64
	//Pascal   float64
	//Bar      float64
	MMH2O float64
}

// Flow is a custom type struct to identify onboard
// pressure sensors
type Flow struct {
	ID       int
	Address  int
	RawValue float64
	Rate     float64
}

// ReadPressure ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Pressure) ReadPressure() {

}

// ReadFlow ... reads raw data from attached sensors to
// Memberane based on its address
func (PS *Flow) ReadFlow() float64 {
	var flowRate float64
	return flowRate
}

// ConvertRawPressure ... converts the raw data to pressure units
// returns Pascal, Bar, mmH2O
func (PS *Pressure) ConvertRawPressure() {

}

// SendToGUI ... sends processed pressure sensor data to GUI
func (PS *Pressure) SendToGUI() {

}
