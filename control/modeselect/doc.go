/*
Package modeselect decides which mode should be used for the ventilator and what related functions whould be
activated depending on the GUI input. There are five conventional modes: volume assist/control;
pressure assist/control; pressure support ventilation;
volume synchronized intermittent mandatory ventilation (SIMV); and pressure SIMV.

	1.	Volume AC
	2.	Pressure AC
	3.	PSV
	4.	SIMV
	5.	Pressure SIMV

For more information look at https://oxfordmedicine.com/view/10.1093/med/9780198784975.001.0001/med-9780198784975-chapter-8.

Volume Assist Control Mode

The volume A/C mode allows two breath types: volume control breaths and volume assist breaths.
Their characteristics are given in the table below.

	Variable	Volume control breath	Volume assist breath
	---------------------------------------------------------
	Triggring	Time			Pressure/flow
	Cycling		Time			Time
	Controlling	Volume			Volume

The volume assist/control mode allows the operator to directly control tidal volume, rate, and Ti for a desired minute volume.
The input variables are:
    a. Tidal volume
    b. Rate
    c. Ti (or I:E ratio or peak flow)
    d. Patient trigger type and sensitivity
    e. PEEP (positive end-expiratory pressure)
    f. FiO2
	g. Flow pattern (possibly).

The triggering window is a defined time slot at late expiration, when the ventilator responds to patient triggering—either by pressure or flow. If the ventilator detects a valid pneumatic signal within the triggering window, it delivers a volume assist breath. If not, it delivers a volume control (time-triggered) breath according to the preset rate. The triggering window is an important technical feature in terms of patient-ventilator synchrony. In a passive patient, all breaths are volume control breaths, and the monitored rate and the set rate are equal. In an active patient, some or all breaths are volume assist breaths, and the monitored rate is higher than the set rate.

The volume A/C mode is intended for ventilated patients who are passive or partially active. It is not a good choice for active patients, especially those with a strong drive, because the patient may not tolerate the inflexible manner of inspiratory gas delivery. Patient-ventilator asynchrony is highly probable.

It is critical to set tidal volume and rate so that the resultant alveolar ventilation matches the patient’s current demand.
Note that the demand may vary during mechanical ventilation. If so, you need to adjust the ventilator settings.



Pressure Assist Control Mode

The pressure A/C mode also allows two breath types: pressure control breaths and pressure assist breaths. Their characteristics are given in the table below.
	Variable	Volume control breath	Volume assist breath
	---------------------------------------------------------
	Triggring	Time			Pressure/flow
	Cycling		Time			Time
	Controlling	Pressure		Pressure

Like the volume A/C mode, the pressure A/C mode has a triggering window, which opens at late expiration. If the ventilator detects a valid pneumatic signal during the triggering window, it delivers a pressure assist breath. If not, it delivers a pressure control (time-triggered) breath according to the set rate. The set Pcontrol applies to both breath types.

In the pressure A/C mode, all breaths are pressure controlled if the ventilated patient is passive, and the monitored rate and the set rate are roughly equal. If the patient is active, some or all breaths are pressure assist breaths, and the monitored rate is typically higher than the set rate.

In the pressure A/C mode, the baseline pressure (PEEP) is constant.

The pressure A/C mode is suitable for passive or partially active patients. It can also be used in active patients with weak respiratory drive, because this mode allows the patient to influence rate, inspiratory flow, and tidal volume. Compared to the volume assist/control mode, pressure assist/control has a considerably lower incidence of patient-ventilator asynchrony. Another advantage of pressure assist/control is that this mode enables the ventilator to compensate for moderate levels of gas leakage.

The perceived disadvantage of this mode is that an operator cannot directly control tidal volume. The resultant tidal volume may be unstable when the patient’s breathing effort and/or respiratory mechanics change. Therefore, you should carefully set the upper and lower limits of the tidal volume alarm.

Pressure support ventilation (PSV) mode

The pressure support ventilation mode allows just one breath type: pressure support breaths.
	Variable	Pressure support breathe
	-----------------------------------------
	Triggring	Pressure/flow
	Cycling		Flow
	Controlling	Pressure

In this mode, an operator sets:
    a. Inspiratory pressure (also known as pressure support)
    b. Patient trigger type and sensitivity
    c. PEEP
    d. FiO2
    e. Flow cycling criteria
    f. Rise time (possibly).

The pressure support ventilation mode is indicated for active patients only. It is the most comfortable mode for this patient population, because they can influence the actual rate, inspiratory time, inspiratory flow, and tidal volume. Obviously, it is contraindicated for the passive patients. Apnoea (backup) ventilation should be activated in this mode. This mode enables the ventilator to adequately compensate for moderate levels of gas leakage.

In pressure support ventilation, the baseline pressure (PEEP) is constant.

You may notice that in some ventilators this mode is called CPAP + PSV. CPAP stands for continuous positive airway pressure. The patient breathes unsupported at an elevated baseline pressure. PSV stands for pressure support ventilation. The patient’s spontaneous breaths are mechanically supported. In this mode, both CPAP and PSV can be realized by changing the pressure support setting.

Spontaneous breaths

If pressure support is set to zero or close to zero, the patient has to do all the required work of breathing to satisfy the ventilatory demand. Unsupported mechanical breaths are spontaneous breaths, often abbreviated as Spont. Spontaneous breaths in ventilated patients are a great challenge, and should be used only in patients who are stable and in good clinical condition. Their typical application is for weaning trials, also known as spontaneous breathing trials.

Typically the patient breathes spontaneously at a moderate positive baseline pressure (PEEP).

Pressure-supported breaths

If you set pressure support to 10 cmH2O or higher, the patient’s spontaneous breaths are pressure supported. In this case, the ventilator takes over a significant part of the required work of breathing.

Pressure-supported breathing is indicated when the ventilated patient is active, but their own efforts are inadequate to meet their required ventilatory demand.

By design, a ventilator in the PSV mode delivers a mechanical breath only when it is pressure or flow triggered. Obviously, this is clinically dangerous because an active patient can stop breathing activity (apnoea) at any time for various clinical reasons. To prevent this potential risk, it is strongly recommended that you activate a protective mechanism called apnoea backup or apnoea ventilation when ventilating in pressure support mode.
(More info at https://oxfordmedicine.com/view/10.1093/med/9780198784975.001.0001/med-9780198784975-chapter-10#med-9780198784975-chapter-10-div1-41)

Volume SIMV mode

The volume SIMV mode allows three breath types: volume control breaths, volume assist breaths, and pressure support breaths.

	Variable	Volume control breath	Volume assist breath	Pressure support breathe
	-----------------------------------------------------------------------------------------
	Triggring	Time			Pressure/flow		Pressure/flow
	Cycling		Time			Time			Flow
	Controlling	Volume			Volume			Pressure

In this mode, an operator sets:

    a. Tidal volume
    b. Rate (also known as SIMV rate)
    c. Ti (or I:E)
    d. Psupport
    e. Patient trigger type and sensitivity
    f. Flow cycling
    g. Rise time (possibly)
    h. PEEP
    i. FiO2.


Volume control breaths are defined by control settings (a), (b), and (c). Volume assist breaths are defined by control settings (a), (c), and (e). Pressure support breaths are defined by control settings (d), (e), (f), and (g).

In the volume SIMV mode, the ventilator delivers volume control breaths at the set SIMV rate. However, if the ventilator detects a valid pressure or flow trigger signal within the triggering window, it delivers a volume assist breath instead. The patient is allowed to breathe spontaneously, with or without pressure support, between two consecutive volume control or assist breaths
*/
package modeselect
