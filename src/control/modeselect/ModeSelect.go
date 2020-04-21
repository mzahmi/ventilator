package modeselect

// VentMode is a custom type struct to identify the mode
// of ventilation required
type VentMode struct {
	ID   int
	Mode string
}

// ReadFromGUI reads input from the GUI to select the
// required Mode
func (M *VentMode) ReadFromGUI() {

}
