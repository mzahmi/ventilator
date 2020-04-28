package temp_sens

import (
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

// TODO: pass in the I2C address as a function argument
// TMP112 chip: if A0 is tied high then address is 0x49
// TMP112 chip: if A0 is tied low then address is 0x48

// const tempSenseI2cAddr = 0x49
const tempRegAddr = 0x00

// note the capital letter for the beginning of the function name means we are exporting
// for use outside the package.  Any functions, variables or consts that begin with
// a lower case character are private to this package

func GetTemperature(i2cAddr uint16) (int8, error) {
	var temp int8

	// init the drivers on the Linux host
	if _, err := host.Init(); err != nil {
		return -127, errors.New("periph.io host.Init() failed")
	}

	// ioexp chip pins the SCL line low if it is in reset.  So even though its not related
	// to the temp sensor, need to drive reset high.
	ioexpRstn := gpioreg.ByName("GPIO24")
	if ioexpRstn == nil {
		return -127, errors.New("failed to init ioexp rst gpio")
	}
	ioexpRstn.Out(gpio.High)

	// Open a handle to the I²C bus:
	bus, err := i2creg.Open("1")
	if err != nil {
		return -127, errors.New("failed to get access to the I2C device")
	}

	// when function exits it will close the bus for you
	defer bus.Close()

	//create an I2C device for the temperature sensor
	tempSensor := &i2c.Dev{Addr: i2cAddr, Bus: bus}

	//set address pointer to temperature register

	// length of write and read slices determine how long the write and reads will be
	// note the 1 byte write length means we only transmit one byte
	// note the zero length read slice means we don't read any bytes
	write := []byte{tempRegAddr}
	read := make([]byte, 0)
	if err := tempSensor.Tx(write, read); err != nil {
		return -127, errors.New("no reply on I2C buss")
	}

	// note the zero length write slice means we won't transmit any bytes
	// note the 2 byte length read slice means we will read two bytes
	write = make([]byte, 0)
	read = make([]byte, 2)

	if err := tempSensor.Tx(write, read); err != nil {
		return -127, errors.New("no reply on I2C buss")
	}

	//throwing away LSBs and truncating to 1 degree C resolution
	temp = int8(read[0])

	return temp, nil
}
