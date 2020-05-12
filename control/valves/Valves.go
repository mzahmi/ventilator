package valves

import (
	// "fmt"

	"github.com/mzahmi/ventilator/control/ioexp"
)

//prop valve inhalation const
const Valve_Prop_Inhalation_Set_DAC_ID = 1
const Valve_Prop_Inhalation_Set_DAC_Chan = 0
const Valve_Prop_Inhalation_Actual_DAC_ID = 1
const Valve_Prop_Inhalation_Actual_DAC_Chan = 1

//Inhalation Valve constants
const Valve_Inhalation_Name = "Inhalation Valve"
const Valve_Inhalation_DAC_ID = 1
const Valve_Inhalation_DAC_Chan = 0

//Exhalation Valve constants
const Valve_Exhalation_Name = "Exhalation Valve"
const Valve_Exhalation_DAC_ID = 1
const Valve_Exhalation_DAC_Chan = 1

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
	Name       string
	SetDacID   uint8
	SetDacChan uint8
	Percent    float64
}

//Solenoids

//MIns ... Inhalation valve
var MIns = SolenValve{
	Name:    Valve_Inhalation_Name,
	State:   false,
	PinMask: ioexp.Solenoid0,
} //normally closed

//MExp ... Exhalation valve
var MExp = SolenValve{
	Name:    Valve_Exhalation_Name,
	State:   true,
	PinMask: ioexp.Solenoid1,
} //normally open

//Proportional Valves

//InProp ... Inhalation proportional valve
var InProp = PropValve{
	Name:       Valve_Inhalation_Name,
	SetDacID:   Valve_Prop_Inhalation_Set_DAC_ID,
	SetDacChan: Valve_Prop_Inhalation_Set_DAC_Chan,
	Percent:    0,
}

//SolenCmd switchs on and off the solenoids
func (valve *SolenValve) SolenCmd(cmd string) {

	switch cmd {
	case "open":
		valve.State = true
		ioexp.WritePin(valve.PinMask, valve.State)
		// fmt.Printf("Valve (%v) has opened\n", valve.Name)
	case "close":
		valve.State = false
		ioexp.WritePin(valve.PinMask, valve.State)
		// fmt.Printf("Valve (%v) has closed\n", valve.Name)

	}

}
