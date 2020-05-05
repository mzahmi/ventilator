//Package alarms ...
package alarms

import (
	"errors"
	"fmt"
	"log"
	"time"

	"vent/control/sensors"
	"vent/pkg/rpigpio"
)

var AlarmReset bool

/* TidalVolumeAlarms sets the upper and lower limits of the tidal volume alarms based on the operator input

High and low tidal volume alarms are based on breath-by-breath tidal volume monitoring.

Tidal volume alarms are mainly relevant in pressure modes where tidal volume varies.
Volume modes generally assure a tidal volume, so that tidal volume alarms should not
be activated under normal conditions of use. However, the tidal volume can sharply
decrease if the peak pressure reaches the limit of the high peak pressure alarm,
causing premature cycling.

Monitored tidal volume can be either inspiratory tidal volume or expiratory tidal volume.
Inspiratory tidal volume is the maximum gas volume that the patient can receive,
while expiratory tidal volume is the minimum gas volume that the patient can receive.
Typically, tidal volume alarms are based on expiratory tidal volume.
Recommended setting for adults:
Upper limit:
	◆ For a passive adult patient, 100 to 150 ml greater than the expected tidal volume
	◆ For an active patient, 50% greater than the expected tidal volume
Lower Limit:
	◆ For a passive adult patient, 100 to 150 ml less than the expected tidal volume
	◆ For an active patient, 50% less than the expected tidal volume
*/
func TidalVolumeAlarms(UpperLimit, LowerLimit float32) error {
	if sensors.FIns.ReadFlow() >= UpperLimit {
		msg := "High tidal volume"
		HighAlert(msg)
		return errors.New(msg)
	} else if sensors.FExp.ReadFlow() <= LowerLimit {
		msg := "Low tidal volume"
		LowAlert(msg)
		return errors.New(msg)
	} else {
		return nil
	}
}

/* AirwayPressureAlarms sets the upper and lower limits of the Peak airway pressure based on the operator input

Recommended setting:
Upper Limit:
	◆ For a passive patient, 10 cmH2O above the expected peak pressure
	◆ For an active patient, 15 cmH2O above the expected peak pressure
Lower Limit:
	◆ For a passive patient, 5 cmH2O below the expected peak pressure
	◆ For an active patient, 5 to 10 cmH2O below the expected peak pressure
*/
func AirwayPressureAlarms(UpperLimit, LowerLimit float32) error {
	if sensors.PIns.ReadPressure() >= UpperLimit {
		msg := "Airway Pressure high"
		HighAlert(msg)
		return errors.New(msg)
	} else if sensors.PIns.ReadPressure() <= LowerLimit {
		msg := "Airway Pressure low"
		LowAlert(msg)
		return errors.New(msg)
	} else {
		return nil
	}
}

/* ExpiratoryMinuteVolumeAlarms

Recommended setting for adults:
Upper limit:
	◆ For a passive patient, 20% greater than the expected minute volume
	◆ For an active patient, 50% greater than the expected minute volume
Lower Limit:
	◆ For a passive patient, 20% less than the expected minute volume
	◆ For an active patient, 50% less than the expected minute volume
*/
func ExpiratoryMinuteVolumeAlarms(UpperLimit, LowerLimit float32) error {
	if sensors.FExp.ReadFlow() >= UpperLimit {
		msg := "High minute volume"
		HighAlert(msg)
		return errors.New(msg)
	} else if sensors.FExp.ReadFlow() <= LowerLimit {
		msg := "Low minute volume"
		LowAlert(msg)
		return errors.New(msg)
	} else {
		return nil
	}
}

/* RespiratoryRateAlarms

High and low respiratory rate alarms are based on the monitored total rate of all
valid mechanical breaths. Respiratory rate directly affects minute volume.

A mechanical breath can be triggered in two ways: time triggering and patient
triggering (pressure of flow). Time triggering is reliable and rigid, while
patient triggering is not 100% reliable. Missed triggering and auto-triggering
are possible.

Recommended setting for adults:
Upper limit:
	◆ For a passive patient, 10 breaths per minute greater than the expected total rate
	◆ For an active patient, 15 breaths per minute greater than the expected total rate
Lower Limit:
	◆ For a passive patient, 10 breaths per minute less than the expected total rate
	◆ For an active patient, 15 breaths per minute less than the expected total rate
*/
func RespiratoryRateAlarms(UpperLimit, LowerLimit float32) error {
	if sensors.FExp.ReadFlow() >= UpperLimit {
		msg := "High Rate"
		HighAlert(msg)
		return errors.New(msg)
	} else if sensors.FExp.ReadFlow() <= LowerLimit {
		msg := "Low Rate"
		LowAlert(msg)
		return errors.New(msg)
	} else {
		return nil
	}
}

/*
 Alarm priorities
*/

/* HighAlert is a High alarm priority

The consequence may be serious injury or death
causes:
	◆ Electrical power or gas failure
	◆ Minute volume too low
	◆ Apnoea
	◆ Airway disconnection

Alarm message on red background

A series of 5 beeps in this sequence, repeated: ▯▯▯_▯▯____▯▯▯_▯▯ */
func HighAlert(msg string) {
	fmt.Println(msg)
	tm := 400 * time.Millisecond
	ts := 3000 * time.Millisecond
	td := 1000 * time.Millisecond
	for i := 1; !AlarmReset; i++ {
		err := rpigpio.BeepOn()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(tm)
		err = rpigpio.BeepOff()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(tm)
		if i%3 == 0 {
			time.Sleep(td)
		}
		if i%5 == 0 {
			time.Sleep(ts)
			i = 0
		}
	}
}

/* MediumAlert is a medium alarm priority

The consequence may be serious if the abnormality persists
causes:
	◆ High total rate
	◆ Inappropriate PEEP/CPAP
	◆ Inappropriate FiO2

Alarm message on yellow background

A series of 3 beeps in this sequence, repeated: ▯▯▯____▯▯▯*/
func MediumAlert(msg string) {
	fmt.Println(msg)
	tm := 400 * time.Millisecond
	ts := 3000 * time.Millisecond
	for i := 1; !AlarmReset; i++ {
		err := rpigpio.BeepOn()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(tm)
		err = rpigpio.BeepOff()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(tm)
		if i%3 == 0 {
			time.Sleep(ts)
			i = 0
		}
	}
}

/* LowAlert is a low alarm priority

The consequence may be moderate if the abnormality persists
causes:
	◆ Compliance/resistance change
	◆ High tidal volume

Alarm message on yellow background

A series of 2 beeps, not repeated: ▯▯*/
func LowAlert(msg string) {
	fmt.Println(msg)
	tm := 400 * time.Millisecond
	err := rpigpio.BeepOn()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(tm)
	err = rpigpio.BeepOff()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(tm)
}
