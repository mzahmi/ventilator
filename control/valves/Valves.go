package valves

import (
	"fmt"
	"vent/pkg/dac"
)

//SolenValve is a custom type struct which contains
//Solenoid ID, Address and State
type SolenValve struct {
	Name    string
	Address string
	state   bool
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
	case "open":
		valve.state = true
		fmt.Printf("Valve (%v) has opened\n", valve.Name)
	case "close":
		valve.state = false
		fmt.Printf("Valve (%v) has closed\n", valve.Name)

	}

}

//IncrementValve adjusts the proportionality of the proportional valve
func (valve *PropValve) IncrementValve(percent float64) {
	valve.Percent = percent
	dac.WriteDac(valve.DacID, valve.DacChan, valve.Percent/10)
	fmt.Printf("Valve (%s) opening has been set to %v\n", valve.Name, valve.Percent)

}
