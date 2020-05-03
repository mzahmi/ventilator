/*
The sensors package declares the sensors onboard the ventilator system as type struct (Pressure and Flow) which indicates
the two types of sensors that are in operation.

The system contains four sensors; two pressure sensors and two flow sensors. They are distributed over the evenly on the
inhilation side and expiration side respictively.

These sensors require hardware calibration to provide the appropriate readings required by the system whether a pressure
reading or a flow reading. These readings can be acquired by two method sets ReadPressure() or ReadFlow().

This package depends on the adc package which is the analog to dc conversion from the membrane board attached to
a raspberry pi 4. To allow us to convert the analog output signal from the onborad sensors into a dc input signal
that can be read by the pi. Refer to adc package for further details.
*/

package sensors
