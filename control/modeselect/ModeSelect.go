package modeselect

import (
	"fmt"
	"github.com/mzahmi/ventilator/control/sensors"
	"time"
)

// VentMode is a custom type struct to identify the mode
// of ventilation required
type VentMode struct {
	ID   int
	Mode string
}

// VolumeAC is a custom type struct that defines the required
// variables for Volume Assist Control ventilation mode
// Trigger is chosen, cycling is time, controlling is volume
type VolumeAC struct {
	TidalVolume float64
	Rate        float64
	Ti          float64
	IR          float64
	ER          float64
	PeakFlow    float64
	PEEP        float64
	FiO2        int
	Trigger     string
	TrigSense   float64
}

// Triggering is a method set for Volume AC mode which returns
// trigger value based three options chosen by operator
// default is Time triggering returns 0
func (VAC *VolumeAC) Triggering() float64 {
	switch VAC.Trigger {
	case "Pressure":
		return VAC.PEEP + VAC.TrigSense
	case "Flow":
		return VAC.TrigSense
	default:
		return 0
	}
}

// Cycling is a method set for Volume AC mode which returns
// the exhalation time Te
func (VAC *VolumeAC) Cycling() float64 {
	var BCT float64 = 60.0 / VAC.Rate
	switch {
	case VAC.Ti != 0:
		return BCT - VAC.Ti
	case VAC.IR != 0:
		VAC.Ti = BCT * VAC.IR / (VAC.IR + VAC.ER)
		return BCT - VAC.Ti
	case VAC.PeakFlow != 0:
		VAC.Ti = VAC.TidalVolume / VAC.PeakFlow
		return BCT - VAC.Ti
	default:
		return BCT - VAC.Ti
	}
}

// Control is a method set for Volume AC mode which controls the vent
func (VAC *VolumeAC) Control() {
	if VAC.Ti != 0 {
		VAC.PeakFlow = VAC.TidalVolume / VAC.Ti
	} else {
		VAC.Ti = VAC.TidalVolume / VAC.PeakFlow
	}
	trig := VAC.Triggering()
	te := VAC.Cycling()
	ti := VAC.Ti
	P1 := sensors.Pressure{
		ID:       1,
		Address:  0,
		RawValue: 255,
		MMH2O:    3,
	}

	go func() {
		for {
			//Turn on the inhalation from reading F1 or P1
			if trig == 0 {
				fmt.Println("Inhale on time trig")
				time.Sleep(time.Second * time.Duration(ti))
				fmt.Println("Exhale on time trig")
				time.Sleep(time.Second * time.Duration(te))
			} else if trig < 0 {
				//reading from P1 trig on calculated value
				if P1.MMH2O == trig {
					fmt.Println("Inhael on P trig")
				}

				fmt.Println("Exhale on P trig")
			}

		}
	}()

}

// ReadFromGUI reads input from the GUI to select the
// required Mode
func (M *VentMode) ReadFromGUI(mode string) {
	switch mode {
	case "Volume A/C":
		fmt.Println("Volume Assisted Control Mode selected")
	case "VC":
		fmt.Println("Volume Control Mode selected")
	case "PC":
		fmt.Println("Pressure Control Mode selected")
	default:
		fmt.Println("No Ventilator Mode selected")
		return
	}
}
