package valves

import "fmt"

//SolenValve is a custom type struct which contains
//Solenoid ID, Address and State
type SolenValve struct {
	Address int
	ID      string
	state   bool
}

//PropValve is a custom type struct which contains
//proportional valve ID, Address and Percentage
type PropValve struct {
	Address int
	ID      string
	Percent int
}

//SolenCmd switchs on and off the solenoids
func (valve *SolenValve) SolenCmd(cmd string) {

	switch cmd {
	case "open":
		valve.state = true
		fmt.Printf("Valve (%v) has opened\n", valve.ID)
	case "close":
		valve.state = false
		fmt.Printf("Valve (%v) has closed\n", valve.ID)

	}

}

//IncrementValve adjusts the proportionality of the proportional valve
func (valve *PropValve) IncrementValve(percent int) {

	valve.Percent = percent
	fmt.Printf("Valve (%v) opening has been set to %v\n", valve.ID, valve.Percent)

}
