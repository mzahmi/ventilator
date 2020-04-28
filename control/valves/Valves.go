package valves

import (
	"fmt"
	"github.com/mzahmi/ventilator/control/dac"
)

//SolenValve is a custom type struct which contains
//Solenoid ID, Address and State
type SolenValve struct {
	ID      int
	Address string
	State   bool
	PinMask uint32
	PinVal  bool
}

//PropValve is a custom type struct which contains
//proportional valve ID, Address and Percentage
type PropValve struct {
	Name    string
	ID      int
	DacChan uint8
	DacID   uint8
	Percent float64
}

//SolenCmd switchs on and off the solenoids
func (valve *SolenValve) SolenCmd(cmd string) {

	switch cmd {
	case "open":
		valve.State = true
		fmt.Printf("Valve (%v) has opened\n", valve.ID)
	case "close":
		valve.State = false
		fmt.Printf("Valve (%v) has closed\n", valve.ID)

	}

}

//IncrementValve adjusts the proportionality of the proportional valve
func (valve *PropValve) IncrementValve(percent float64) {
	valve.Percent = percent
	dac.WriteDac(valve.ID, valve.DacChan, valve.Percent/10)
	fmt.Printf("Valve (%s) opening has been set to %v\n", valve.Name, valve.Percent)

}
