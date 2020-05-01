package rpigpio

import (
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func BeepOn() error {
	err := gpioHigh("GPIO12")
	if err != nil {
		return err
	}
	return nil
}

func BeepOff() error {
	err := gpioLow("GPIO12")
	if err != nil {
		return err
	}
	return nil
}

func gpioHigh(gpio_name string) error {
	// init the drivers on the Linux host
	if _, err := host.Init(); err != nil {
		return errors.New("periph.io host.Init() failed")
	}
	gp := gpioreg.ByName(gpio_name)
	if gp == nil {
		return errors.New("failed to access GPIO")
	}
	gp.Out(gpio.High)
	return nil
}

func gpioLow(gpio_name string) error {
	// init the drivers on the Linux host
	if _, err := host.Init(); err != nil {
		return errors.New("periph.io host.Init() failed")
	}
	gp := gpioreg.ByName(gpio_name)
	if gp == nil {
		return errors.New("failed to access GPIO")
	}
	gp.Out(gpio.Low)
	return nil
}

// returning 6 pointers to the buttons and an error
func InitButtons() (gpio.PinIO, gpio.PinIO, gpio.PinIO, gpio.PinIO, gpio.PinIO, gpio.PinIO, error) {

	// btn1 -> GPIO19
	// btn2 -> GPIO13
	// btn3 -> GPIO6
	// btn4 -> GPIO5
	// btn5 -> GPIO22
	// btn6 -> GPIO27

	// yes its super verbose, but if anything goes wrong it will tell you exactly what its ticked off about

	btn1 := gpioreg.ByName("GPIO19")
	if btn1 == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("btn1 GPIO init failed")
	}
	btn2 := gpioreg.ByName("GPIO13")
	if btn1 == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("btn2 GPIO init failed")
	}
	btn3 := gpioreg.ByName("GPIO6")
	if btn1 == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("btn3 GPIO init failed")
	}
	btn4 := gpioreg.ByName("GPIO5")
	if btn1 == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("btn4 GPIO init failed")
	}
	btn5 := gpioreg.ByName("GPIO22")
	if btn1 == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("btn5 GPIO init failed")
	}
	btn6 := gpioreg.ByName("GPIO27")
	if btn1 == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("btn6 GPIO init failed")
	}
	return btn1, btn2, btn3, btn4, btn5, btn6, nil
}
