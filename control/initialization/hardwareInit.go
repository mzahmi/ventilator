package initialization

import (
	"fmt"
	"time"

	"github.com/mzahmi/ventilator/control/adc"
	"github.com/mzahmi/ventilator/control/dac"
	"github.com/mzahmi/ventilator/control/ioexp"
	"github.com/mzahmi/ventilator/control/rpigpio"
)

//InitHardware ... this function should be called at the beginning of main.
//It will initialize all the hardware and check for errors
func HardwareInit() {
	fmt.Println("Beginning hardware initialization")

	//I2C init
	initI2C()
	//init ADC
	initADC()
	//init DAC
	initDAC()

	fmt.Println("End of hardware initialization")

}

//initI2C ... this function initializes I2C and checks for errors
//initI2C ... this function initializes I2C and checks for errors
func initI2C() {
	fmt.Println("Starting I2C init")
	err := ioexp.InitChip()
	if err != nil {
		fmt.Println(err)
		return
	}
	//Beep test
	fmt.Println("beep called")
	for ii := 0; ii < 3; ii++ {

		err = rpigpio.BeepOn()
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(50 * time.Millisecond)
		err = rpigpio.BeepOff()
		if err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(50 * time.Millisecond)
	}

	//testing LEDs
	const blinkTime = 200
	fmt.Println("Blinking LEDs")
	for ii := 0; ii < 2; ii++ {

		fmt.Println("Yellow")
		err := ioexp.WritePin(ioexp.YellowLed, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(blinkTime * time.Millisecond)
		err = ioexp.WritePin(ioexp.YellowLed, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Red")
		err = ioexp.WritePin(ioexp.RedLed, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(blinkTime * time.Millisecond)
		err = ioexp.WritePin(ioexp.RedLed, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Green")
		err = ioexp.WritePin(ioexp.GreenLed, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(blinkTime * time.Millisecond)
		err = ioexp.WritePin(ioexp.GreenLed, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Blue")
		err = ioexp.WritePin(ioexp.BlueLed, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(blinkTime * time.Millisecond)
		err = ioexp.WritePin(ioexp.BlueLed, false)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//initADC ... this function initializes ADC and checks for errors
func initADC() {

	fmt.Println("readAdc called")

	_, err := adc.ReadADC(1)
	if err != nil {
		fmt.Println(err)
		return
	}

}

//initDAC ... this function initializes DAC and checks for errors
func initDAC() {

	fmt.Println("dacsZero called")
	err := dac.DacsAllZeroOut()
	if err != nil {
		fmt.Println(err)
	}

}
