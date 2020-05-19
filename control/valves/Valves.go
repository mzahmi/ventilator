package valves

import (
	// "fmt"

	"github.com/mzahmi/ventilator/control/ioexp"
	"github.com/mzahmi/ventilator/pkg/dac"
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
	Enable     uint32
	Percent    float64
}

//AnyValve interfaces all types of valves involved in the system
type AnyValve interface {
	Close()
}

//Solenoids
//MV ... Main Valve
var MV = SolenValve{
	Name:    Valve_Inhalation_Name,
	State:   false,
	PinMask: ioexp.Solenoid0,
}

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
	Enable:     ioexp.Solenoid2,
	Percent:    0,
}

//InProp ... Inhalation proportional valve
var ExProp = PropValve{
	Name:       Valve_Inhalation_Name,
	SetDacID:   Valve_Prop_Inhalation_Set_DAC_ID,
	SetDacChan: Valve_Prop_Inhalation_Set_DAC_Chan,
	Enable:     0,
	Percent:    0,
}

//Open opens the solen valve
func (valve *SolenValve) Open() {
	valve.State = true
	ioexp.WritePin(valve.PinMask, valve.State)
}

//Close closes the solen valve
func (valve *SolenValve) Close() {
	valve.State = false
	ioexp.WritePin(valve.PinMask, valve.State)
}

//IncrementValve controls the opening of Prop valves
func (valve *PropValve) IncrementValve(actuate float64) {
	ioexp.WritePin(valve.Enable, true)
	dac.WriteDac(valve.SetDacID, valve.SetDacChan, actuate)
}

//Close closes prop valve
func (valve *PropValve) Close() {
	ioexp.WritePin(valve.Enable, false)
	valve.IncrementValve(0)
}

func CloseAllValves(v ...AnyValve) {
	for _, s := range v {
		s.Close()
	}
}
