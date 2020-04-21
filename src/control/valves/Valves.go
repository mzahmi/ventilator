package valves

import "fmt"

//state is true if valve is open
type SolenValve struct {
	Address int
	ID      string
	state   bool
}

type PropValve struct {
	Address int
	ID      string
	Percent int
}

func (valve *SolenValve) SolenCmd(cmd string) {

	switch cmd {
	case "open":
		valve.state = true
		fmt.Println("Valve (%v)", valve.ID, " has opened")
	case "close":
		valve.state = false
		fmt.Println("Valve (%v)", valve.ID, " has opened")

	}

}

func (valve *PropValve) IncrementValve(percent int) {

	valve.Percent = percent
	fmt.Println("Valve (%v)", valve.ID, " opening has been set to %v", valve.ID)

}
