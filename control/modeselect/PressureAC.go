package modeselect

import (
	"fmt"
	"log"
	"time"

	"github.com/mzahmi/ventilator/alarms"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/control/valves"
)

/* PressureAC has a triggering window, which opens at late expiration.
If the ventilator detects a valid pneumatic signal during the triggering window,
it delivers a pressure assist breath. If not, it delivers a pressure control
(time-triggered) breath according to the set rate. The set Pcontrol applies to
both breath types (Pressure Control Breath and Preassure Assist Breath).

In the pressure A/C mode, all breaths are pressure controlled if the ventilated
patient is passive, and the monitored rate and the set rate are roughly equal.
If the patient is active, some or all breaths are pressure assist breaths,
and the monitored rate is typically higher than the set rate.

In the pressure A/C mode, the baseline pressure (PEEP) is constant.

The pressure A/C mode is suitable for passive or partially active patients.
It can also be used in active patients with weak respiratory drive,
because this mode allows the patient to influence rate, inspiratory flow,
and tidal volume. Compared to the volume assist/control mode,
pressure assist/control has a considerably lower incidence of patient-ventilator
asynchrony. Another advantage of pressure assist/control is that this mode enables
the ventilator to compensate for moderate levels of gas leakage.

The perceived disadvantage of this mode is that an operator cannot directly
control tidal volume. The resultant tidal volume may be unstable when the patient’s
breathing effort and/or respiratory mechanics change. Therefore, you should
carefully set the upper and lower limits of the tidal volume alarm.*/
func PressureAC(UI *UserInput) {
	switch UI.BreathType {
	case "Control":
		PressureControl(UI)
	case "Assist":
		PressureAssist(UI)
	default:
		fmt.Println("Enter valid breath type")
	}
}

// PressureControl a submode of PressureAC.
// 	Triggering:	Time
// 	Cycling: 	Time
// 	Control: 	Pressure
func PressureControl(UI *UserInput) {
	//initiate Pressure PID based on readings from PIns
	PressurePID := NewPIDController(0.5, 0.5, 0.5)         // takes in P, I, and D values to be set trial and error
	PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets the PID setpoint to inspiratory pressure

	//control loop; it loops unitll Exit bool is set to false
	for !Exit {
		//Open main valve MIns controlled by pressure sensor PIns
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
			//check for alarms
			valves.InProp.IncrementValve(PressurePID.Update(float64(sensors.PIns.ReadPressure())))
		}

		//Close main valve InProp
		valves.InProp.IncrementValve(0)

		//Open main valve MExp controlled by pressure sensor PExp
		for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
			// check for alarms
			if sensors.PExp.ReadPressure() <= UI.PEEP {
				break
			}
			valves.ExProp.IncrementValve(100)
		}
		//Close main valve ExProp
		valves.ExProp.IncrementValve(0)
	}

}

// PressureAssist a submode of PressureAC. The Triggering type is chosen by the operator.
// 	Triggering:	Pressure/Flow
// 	Cycling: 	Time
// 	Control: 	Pressure
func PressureAssist(UI *UserInput) {
	//initiate Pressure PID based on readings from PIns
	PressurePID := NewPIDController(0.5, 0.5, 0.5)         // takes in P, I, and D values to be set trial and error
	PressurePID.setpoint = float64(UI.InspiratoryPressure) // Sets PID setpoint to Inspiratory pressure

	//Check trigger type
	switch UI.PatientTriggerType {
	case "Pressure":
		//Calculate trigger threshhold with PEEP and sensitivity
		PTrigger := UI.PEEP + UI.PressureTrigSense

		//control loop; it loops unitll Exit bool is set to false
		for !Exit {
			//check if trigger is true
			if sensors.PIns.ReadPressure() <= PTrigger {

				//Open main valve InProp controlled by pressure sensor PIns and check tidal volume alarms
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
					//check for tidal volume alarms
					err := alarms.TidalVolumeAlarms(UI.UpperLimitVT, UI.LowerLimitVt)
					if err != nil {
						log.Println(err)
						//or we can use panic
						//panic(err)
						break
					}
					valves.InProp.IncrementValve(PressurePID.Update(float64(sensors.PIns.ReadPressure())))
				}
				//Close main valve InProp
				valves.InProp.IncrementValve(0) // closes the valve

				//Open main valve ExProp and check for PEEP value and tidal volume alarms
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
					//check for tidal volume alarms
					err := alarms.TidalVolumeAlarms(UI.UpperLimitVT, UI.LowerLimitVt)
					if err != nil {
						log.Println(err)
						//or we can use panic
						//panic(err)
						break
					}
					if sensors.PExp.ReadPressure() <= UI.PEEP {
						break
					}
					valves.ExProp.IncrementValve(100)
				}
				//Close main valve ExProp
				valves.ExProp.IncrementValve(0) // closes the valve
			}
		}
	case "Flow":
		//Calculate trigger threshhold with flow trig sensitivity
		FTrigger := UI.FlowTrigSense
		//control loop; it loops unitll Exit bool is set to false
		for !Exit {
			//check if trigger is true
			if sensors.FIns.ReadFlow() >= FTrigger {
				//Open main valve InProp controlled by pressure sensor PIns
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Ti*1000) * time.Millisecond); {
					//check for tidal volume alarms
					err := alarms.TidalVolumeAlarms(UI.UpperLimitVT, UI.LowerLimitVt)
					if err != nil {
						log.Println(err)
						//or we can use panic
						//panic(err)
						break
					}
					valves.InProp.IncrementValve(PressurePID.Update(float64(sensors.PIns.ReadPressure())))
				}

				//Close main valve InProp
				valves.InProp.IncrementValve(0) // closes the valve

				//Open main valve ExProp and check for PEEP value and tidal volume alarms
				for start := time.Now(); time.Since(start) < (time.Duration(UI.Te*1000) * time.Millisecond); {
					//check for tidal volume alarms
					err := alarms.TidalVolumeAlarms(UI.UpperLimitVT, UI.LowerLimitVt)
					if err != nil {
						log.Println(err)
						//or we can use panic
						//panic(err)
						break
					}
					if sensors.PExp.ReadPressure() <= UI.PEEP {
						break
					}
					valves.ExProp.IncrementValve(100)
				}
				//Close main valve MExp
				valves.ExProp.IncrementValve(0) // closes the valve
			}
		}
	}
}
