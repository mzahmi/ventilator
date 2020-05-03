/*
Ventilator alarms are based on a fairly simple wish: we clinicians want the ventilator to alert us whenever a ventilated patient faces potential danger associated with the mechanical ventilation.

Accomplishing this involves working out the following:

1. The alarm conditions—those under which a ventilated patient may be uncomfortable, injured, or even die;
2. How an alarm condition will be detected, including (a) which monitored parameters are used, and (b) what the normal ranges (non-alarm zones) for those parameters are;
3. Timing—when the alarm should be activated if the alarm condition is detected (i.e. immediately or with a slight delay);
4. The alarm message to be displayed.
As an example, let’s take a look at the high peak pressure alarm, which we call simply the ‘high pressure’ alarm.

It is well known that excessive airway pressure can cause barotrauma to the lungs. So we want the ventilator to alert us if the peak pressure is dangerously high. Clearly, the monitored peak airway pressure should be used as the foundation of this alarm. In ventilator alarm design, we also recognize that a threshold is necessary to differentiate between normal and high peak airway pressure. The threshold may differ under various clinical conditions, so it should be set individually by the clinicians in charge. To prevent the alarm from being oversensitive, we decide that the alarm ‘high peak pressure’ should be activated only when the monitored peak pressure exceeds the threshold for two consecutive breaths.

Being so designed, after every breath the ventilator compares the monitored peak pressure with the operator-set threshold. The alarm, with its visual and audible indications, is activated only when the defined alarm condition is detected. Otherwise, the alarm is inactive.

The principle is applicable to all ventilator alarms, simple or complex.

Technical Alarms And Application Alarms

A ventilator can have a number of alarms. They can be roughly divided into two categories: technical alarms and application alarms.

	Technical alarms											Application alarms
	-------------------------------------------------------------------------------------------------------------
	◆ Technical problems of ventilator and 						◆ Leak or occlusion of ventilator system in which
	accessories (design, production, and maintenance)			ventilator and accessories are technically in order
	◆ Electrical supply problems								◆ Problems with patient’s pulmonary system,
	◆ Gas (air and oxygen) supply problems						e.g. pneumothorax
																◆ Improper ventilator setting

Technical Alarms

Technical alarms are related to abnormalities of the ventilator itself, the electrical supply, or the gas (air and oxygen) supplies. Ventilator abnormalities can have various origins, such as problems with ventilator design and development, problems with production, and problems with device maintenance and service.

Application Alarms

At this point, let’s return to some fundamental points about mechanical ventilation. We know that the equipment required for mechanical ventilation is a ventilator system composed of six parts. A ventilator is just one of them. A ventilator system works properly only when all three conditions are satisfied:

1. All parts are present, functioning, compatible, and properly connected.
2. The gas passageway inside the ventilator system is neither leaking nor obstructed.
3. The operation of the ventilator system is adapted appropriately and individually to the patient’s clinical conditions.
Obviously, it is important that we know when a ventilator system is not functioning properly. Application alarms are designed to signal such functional abnormalities. Note that with an application alarm, the ventilator itself is often technically in order.

Application alarms occur far more frequently than technical alarms. However, application alarms may be less understood.

Typically technical personnel (e.g. engineers from ventilator manufacturing firms and hospital technicians) deal with technical alarms. Ventilator operators (clinicians) deal with application alarms, so this chapter will focus on application alarms.
(More info at: https://oxfordmedicine.com/view/10.1093/med/9780198784975.001.0001/med-9780198784975-chapter-12)
*/

package alarms
