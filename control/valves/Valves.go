package valves

import (
	"fmt"

	"github.com/mzahmi/ventilator/control/dac"
	"github.com/mzahmi/ventilator/control/ioexp"
)

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
	Name    string
	DacChan uint8
	DacID   uint8
	Percent float64
}

type Valves interface {
	CloseValve()
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
	Name:    Valve_Inhalation_Name,
	DacID:   Valve_Inhalation_DAC_ID,
	DacChan: Valve_Inhalation_DAC_Chan,
	Percent: 0,
}

//ExProp ... Exhalation proportional valve
var ExProp = PropValve{
	Name:    Valve_Exhalation_Name,
	DacID:   Valve_Exhalation_DAC_ID,
	DacChan: Valve_Exhalation_DAC_Chan,
	Percent: 0,
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

//CloseValve closes solenvalve
func (valve *SolenValve) CloseValve() {
	valve.SolenCmd("Close")
}

//IncrementValve adjusts the proportionality of the proportional valve
func (valve *PropValve) IncrementValve(percent float64) {
	valve.Percent = percent
	dac.WriteDac(valve.DacID, valve.DacChan, valve.Percent/10)
	//fmt.Printf("Valve (%s) opening has been set to %v\n", valve.Name, valve.Percent)

}

// CloseVale closes propvalves
func (valve *PropValve) CloseValve() {
	valve.IncrementValve(0)
}

//CloseAllValves takes in either a SolenValve or a PropValve to close them
func CloseAllValves(v ...Valves) {
	for _, s := range v {
		s.CloseValve()
	}
}
