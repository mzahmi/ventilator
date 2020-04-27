package valves

import "fmt"

//SolenValve is a custom type struct which contains
//Solenoid ID, Address and State
type SolenValve struct {
	ID      string
	Address string
	state   bool
}

//PropValve is a custom type struct which contains
//proportional valve ID, Address and Percentage
type PropValve struct {
	ID      string
	Address string
	Percent float64
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
func (valve *PropValve) IncrementValve(percent float64) {
	valve.Percent = percent
	fmt.Printf("Valve (%s) opening has been set to %v\n", valve.ID, valve.Percent)

}
