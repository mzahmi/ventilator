package valves

import (
	"fmt"
	"github.com/mzahmi/ventilator/control/dac"
	"github.com/mzahmi/ventilator/control/ioexp"
)

//SolenValve is a custom type struct which contains
//Solenoid ID, Address and State
type SolenValve struct {
	Name    string
	State   bool
	PinMask uint32
}

//PropValve is a custom type struct which contains
//proportional valve ID, Address and Percentage
type PropValve struct {
	Name    string
	DacChan uint8
	DacID   uint8
	Percent float64
}

//SolenCmd switchs on and off the solenoids
func (valve *SolenValve) SolenCmd(cmd string) {

	switch cmd {
	case "Open":
		valve.State = true
		ioexp.WritePin(valve.PinMask, valve.State)
		fmt.Printf("Valve (%v) has opened\n", valve.Name)
	case "Close":
		valve.State = false
		ioexp.WritePin(valve.PinMask, valve.State)
		fmt.Printf("Valve (%v) has closed\n", valve.Name)

	}

}

//IncrementValve adjusts the proportionality of the proportional valve
func (valve *PropValve) IncrementValve(percent float64) {
	valve.Percent = percent
	dac.WriteDac(valve.DacID, valve.DacChan, valve.Percent/10)
	fmt.Printf("Valve (%s) opening has been set to %v\n", valve.Name, valve.Percent)

}
