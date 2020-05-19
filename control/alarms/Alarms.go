//Package alarms ...
package alarms

import (
	"errors"
	"runtime"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/mzahmi/ventilator/control/sensors"
	"github.com/mzahmi/ventilator/logger"
	"github.com/mzahmi/ventilator/pkg/rpigpio"
)

var AlarmReset bool // updated from redis client "status"

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
func TidalVolumeAlarms(s *sensors.SensorsReading, mux *sync.Mutex, UpperLimit, LowerLimit float32, logStruct *logger.Logging, client *redis.Client) error {
	mux.Lock()
	trig := s.PressureInput // TODO: needs to be a flow sensor
	mux.Unlock()
	runtime.Gosched()
	if trig >= UpperLimit {
		msg := "High tidal volume"
		info := "Check tidal volume"
		HighAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else if trig <= LowerLimit {
		msg := "Low tidal volume"
		info := "Check tidal volume"
		LowAlert(msg, info, logStruct, client)
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

	air way pressure is PEEP + (pressure insp or pressure support)
*/
func AirwayPressureAlarms(s *sensors.SensorsReading, mux *sync.Mutex, UpperLimit, LowerLimit float32, logStruct *logger.Logging, client *redis.Client) error {
	mux.Lock()
	trig := s.PressureInput
	mux.Unlock()
	runtime.Gosched()
	if trig >= UpperLimit {
		msg := "Airway Pressure high"
		info := "Check for airway obstruction"
		HighAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else if trig <= LowerLimit {
		msg := "Airway Pressure low"
		info := "check for gas leakage"
		LowAlert(msg, info, logStruct, client)
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
func ExpiratoryMinuteVolumeAlarms(s *sensors.SensorsReading, mux *sync.Mutex, UpperLimit, LowerLimit float32, logStruct *logger.Logging, client *redis.Client) error {
	mux.Lock()
	trig := s.PressureOutput // TODO: should be Flowout
	mux.Unlock()
	runtime.Gosched()
	if trig >= UpperLimit {
		msg := "High minute volume"
		info := "check for tidal volume and RR"
		HighAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else if trig <= LowerLimit {
		msg := "Low minute volume"
		info := "check for tidal volume and RR"
		LowAlert(msg, info, logStruct, client)
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
func RespiratoryRateAlarms(UpperLimit, LowerLimit float32, logStruct *logger.Logging, client *redis.Client) error {
	RRM := float32(20.0) // TODO: find a way to monitor the BPM of the patient
	if RRM >= UpperLimit {
		msg := "High Rate"
		info := "check patient breating mechanism"
		HighAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else if RRM <= LowerLimit {
		msg := "Low Rate"
		info := "check patient breating mechanism"
		LowAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else {
		return nil
	}
}

/*
Technical Alarms
A ventilator system requires stable and continuous supplies of pressurized oxygen
and air for its proper functioning. The normal working pressure of both supplies
is typically between 2 or 3 and 6 or 6.5 bars (between 200 or 300 and 600 or 650 kPa,
or between 29 or 43.5 and 87 or 94 psi)
*/

/*
OxygenSupplyAlarm indicates oxygen supply pressure is less than the lower limit of 2 or 3 bars

The oxygen supply pressure is too low because of:
	◆ Unexpected interruption of central oxygen supply
	◆ A significant leak or disconnection in the oxygen supply route, i.e. hose or fitting
	◆ Nearly empty oxygen cylinder

If an air supply is available, mechanical ventilation should continue with air alone
*/
func OxygenSupplyAlarm(s *sensors.SensorsReading, mux *sync.Mutex, LowerO2Press float32, logStruct *logger.Logging, client *redis.Client) error {
	mux.Lock()
	trig := s.PressureInput // TODO: should be an O2 sensor reading
	mux.Unlock()
	runtime.Gosched()
	if trig <= LowerO2Press { // change to oxygen supply sensor reading
		msg := "Low O2 supply"
		info := "check O2 inlet for leakage or low pressure"
		MediumAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else {
		return nil
	}
}

/*
AirSupplyAlarms indicates air supply pressure is less than the lower limit of 2 or 3 bars

The air supply pressure is too low because of:
	◆ Unexpected interruption of central air supply
	◆ A significant leak or disconnection in the air supply route, i.e. hose or fitting
	◆ Nearly empty air cylinder
	◆ Defective air compressor

If an oxygen supply is available, mechanical ventilation should continue with 100% oxygen
*/
func AirSupplyAlarm(s *sensors.SensorsReading, mux *sync.Mutex, LowerAirPress float32, logStruct *logger.Logging, client *redis.Client) error {
	mux.Lock()
	trig := s.PressureInput // TODO: should be a pressure sensor at the entry of the system
	mux.Unlock()
	runtime.Gosched()
	if trig <= LowerAirPress { // change to air supply sensor reading
		msg := "Low Air supply"
		info := "check air inlet for leakage or low pressure"
		MediumAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else {
		return nil
	}
}

/*
AirAndO2SupplyAlarm indicates both air and oxygen supply pressures are less than 2 or 3 bars

Both air and oxygen supply pressures are too low

If both gas supplies fail at the same time, a ventilator system cannot continue to function.
The ventilator automatically switches to the ambient state.
*/
func AirAndO2SupplyAlarm(airerr, o2err error, logStruct *logger.Logging, client *redis.Client) error {
	if (airerr != nil) && (o2err != nil) {
		msg := "Gas supply is low"
		info := "check gas inlet for leakage or low pressure"
		HighAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else {
		return nil
	}
}

/*
FiO2Alarms ndicate that the current inspiratory O2 concentration is outside the set range
High Alram:
	The monitored FiO2 is 5% to 7% above the set FiO2 for a defined duration, e.g. 30 s
Low Alarm:
	The monitored FiO2 is 5% or 7% below the set FiO2 for a defined duration, e.g. 30 s
Common causes:
	◆ Faulty ventilator mixing function
	◆ Faulty oxygen monitoring, e.g. defective or uncalibrated oxygen cell
	◆ Low FiO2 due to use of an oxygen concentrator. Standard oxygen supplies provide pure oxygen. O2 from a concentrator may be as low as 90%
*/

func FiO2Alarms(s *sensors.SensorsReading, mux *sync.Mutex, UpperLimit, LowerLimit float32, logStruct *logger.Logging, client *redis.Client) error {
	mux.Lock()
	trig := s.PressureInput //TODO: read from new O2 sensor at the entry of the system
	mux.Unlock()
	runtime.Gosched()
	if trig >= UpperLimit { // change to oxygen sensor reading
		msg := "FiO2 is High"
		info := "check O2 concentration"
		HighAlert(msg, info, logStruct, client)
		return errors.New(msg)
	} else if trig <= LowerLimit { // change to oxygen sensor reading
		msg := "FiO2 is Low"
		info := "check O2 concentration"
		LowAlert(msg, info, logStruct, client)
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
func HighAlert(title, info string, logStruct *logger.Logging, client *redis.Client) {
	// logArm.Println(msg)
	client.Set("alarm_status", "high", 0).Err()
	client.Set("alarm_title", title, 0).Err()
	client.Set("alarm_text", info, 0).Err()
	logStruct.Alarm(title)
	tm := 400 * time.Millisecond
	ts := 3000 * time.Millisecond
	td := 1000 * time.Millisecond
	for i := 1; !AlarmReset; i++ {
		err := rpigpio.BeepOn()
		check(err, logStruct)
		time.Sleep(tm)
		err = rpigpio.BeepOff()
		check(err, logStruct)
		time.Sleep(tm)
		if i%3 == 0 {
			time.Sleep(td)
		}
		if i%5 == 0 {
			time.Sleep(ts)
			i = 0
		}
		status, err := client.Get("alarm_status").Result()
		check(err, logStruct)
		if status == "none" {
			AlarmReset = true
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
func MediumAlert(title, info string, logStruct *logger.Logging, client *redis.Client) {
	// logArm.Println(msg)
	client.Set("alarm_status", "medium", 0).Err()
	client.Set("alarm_title", title, 0).Err()
	client.Set("alarm_text", info, 0).Err()
	logStruct.Alarm(title)
	tm := 400 * time.Millisecond
	ts := 3000 * time.Millisecond

	for i := 1; !AlarmReset; i++ {
		err := rpigpio.BeepOn()
		check(err, logStruct)
		time.Sleep(tm)
		err = rpigpio.BeepOff()
		check(err, logStruct)
		time.Sleep(tm)
		if i%3 == 0 {
			time.Sleep(ts)
			i = 0
		}
		status, err := client.Get("alarm_status").Result()
		check(err, logStruct)
		if status == "none" {
			AlarmReset = true
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
func LowAlert(title, info string, logStruct *logger.Logging, client *redis.Client) {
	// logArm.Println(msg)
	client.Set("alarm_status", "low", 0).Err()
	client.Set("alarm_title", title, 0).Err()
	client.Set("alarm_text", info, 0).Err()
	logStruct.Alarm(title)
	tm := 400 * time.Millisecond
	err := rpigpio.BeepOn()
	check(err, logStruct)
	time.Sleep(tm)
	err = rpigpio.BeepOff()
	check(err, logStruct)
	time.Sleep(tm)
	status, err := client.Get("alarm_status").Result()
	check(err, logStruct)
	if status == "none" {
		AlarmReset = true
	}
}

// prints out the checked error err
func check(err error, logStruct *logger.Logging) {
	if err != nil {
		logStruct.Err(err)
	}
}
