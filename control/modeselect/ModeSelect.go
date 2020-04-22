package modeselect

import "fmt"

// VentMode is a custom type struct to identify the mode
// of ventilation required
type VentMode struct {
	ID   int
	Mode string
}

// ReadFromGUI reads input from the GUI to select the
// required Mode
func (M *VentMode) ReadFromGUI(mode string) {
	switch mode {
	case "AC":
		fmt.Println("Assisted Control Mode selected")
	case "VC":
		fmt.Println("Volume Control Mode selected")
	case "PC":
		fmt.Println("Pressure Control Mode selected")
	default:
		fmt.Println("No Ventilator Mode selected")
		return
	}
}
