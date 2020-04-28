package i2c_mux

import (
	"errors"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

// all address lines are low (see table one of data sheet for address)
const i2cMuxAddr = 0x70

func SetI2CMux(chanVal uint8) error {

	if chanVal > 7 {
		return errors.New("Channel selected must be between 0 and 7")
	}

	// set bit associated with channel
	chanCode := 1 << chanVal

	// init the drivers on the Linux host
	if _, err := host.Init(); err != nil {
		return errors.New("periph.io host.Init() failed")
	}

	// ioexp chip pins the SCL line low if it is in reset.  So even though its not related
	// to the temp sensor, need to drive reset high.
	ioexpRstn := gpioreg.ByName("GPIO24")
	if ioexpRstn == nil {
		return errors.New("failed to init ioexp rst gpio")
	}
	ioexpRstn.Out(gpio.High)

	i2cMuxRstn := gpioreg.ByName("GPIO16")
	if i2cMuxRstn == nil {
		return errors.New("failed to init i2c mux rst gpio")
	}
	i2cMuxRstn.Out(gpio.High)

	// Open a handle to the I²C bus:
	bus, err := i2creg.Open("1")
	if err != nil {
		return errors.New("failed to get access to the I2C device")
	}

	// when function exits it will close the bus for you
	defer bus.Close()

	//create an I2C device for the temperature sensor
	mux := &i2c.Dev{Addr: i2cMuxAddr, Bus: bus}

	write := []byte{byte(chanCode)}
	read := make([]byte, 0)
	if err := mux.Tx(write, read); err != nil {
		return errors.New("no reply from I2C mux on I2C buss")
	}

	return nil
}
