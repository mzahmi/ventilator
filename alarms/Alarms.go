package alarms

import (
	"errors"

	"github.com/mzahmi/ventilator/control/sensors"
)

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
	◆ For an active patient, 50% less than the expected tidal volume */
func TidalVolumeAlarms(UpperLimit, LowerLimit float32) error {
	if sensors.FIns.ReadFlow() >= UpperLimit {
		return errors.New("Upper limit of tidal volume reached")
	} else if sensors.FExp.ReadFlow() <= LowerLimit {
		return errors.New("Upper limit of tidal volume reached")
	} else {
		return nil
	}
}
