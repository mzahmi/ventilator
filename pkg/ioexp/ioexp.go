// driver for the TCA6424A IO expander chip

package ioexp

import (
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

const ioexpI2cAddr = 0x22
const p0Default = 0x00      // 0000 0001
const p1Default = 0x00      // 0000 0100
const p2Default = 0x00      // 00const
const outPortAddr = 0x84    // auto increment flag on
const configPortAddr = 0x8C // auto increment flag on

// bit masks for all the ports
const YellowLed = (1 << 0)
const RedLed = (1 << 1)
const GreenLed = (1 << 2)
const BlueLed = (1 << 3)
const Solenoid0 = (1 << 4)
const Solenoid1 = (1 << 5)
const Solenoid2 = (1 << 6)
const Solenoid3 = (1 << 7)
const Solenoid4 = (1 << 8)
const Solenoid5 = (1 << 9)
const Solenoid6 = (1 << 10)
const Solenoid7 = (1 << 11)
const Opto1 = (1 << 12)
const Opto2 = (1 << 13)
const Opto3 = (1 << 14)
const Opto4 = (1 << 15)
const P20 = (1 << 16)
const P21 = (1 << 17)
const P22 = (1 << 18)
const P23 = (1 << 19)
const P24 = (1 << 20)
const P25 = (1 << 21)
const P26 = (1 << 22)
const P27 = (1 << 23)

func InitChip() error {
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

	bus, err := i2creg.Open("1")
	if err != nil {
		return errors.New("failed to get access to the I2C device")
	}
	defer bus.Close()
	ioexp := &i2c.Dev{Addr: ioexpI2cAddr, Bus: bus}
	write := []byte{outPortAddr, p0Default, p1Default, p2Default}
	read := make([]byte, 0)
	if err := ioexp.Tx(write, read); err != nil {
		return errors.New("no reply on I2C buss in init a")
	}
	// configure all ports as outputs for now (easier to test hardware)
	write = []byte{configPortAddr, 0x00, 0x00, 0x00}
	read = make([]byte, 0)
	if err := ioexp.Tx(write, read); err != nil {
		return errors.New("no reply on I2C buss in init b")
	}
	return nil
}

// read modify write so you can change one pin at a time.  This only works if all the ports are outputs.
// Will need to do a better driver when it is time to to mix inputs and outputs

func WritePin(pinMask uint32, pinVal bool) error {
	if _, err := host.Init(); err != nil {
		return errors.New("periph.io host.Init() failed")
	}
	bus, err := i2creg.Open("1")
	if err != nil {
		return errors.New("failed to get access to the I2C device")
	}
	defer bus.Close()
	ioexp := &i2c.Dev{Addr: ioexpI2cAddr, Bus: bus}
	write := []byte{outPortAddr}
	read := make([]byte, 3)
	// read the current value of all output registers
	if err := ioexp.Tx(write, read); err != nil {
		return errors.New("no reply on I2C buss")
	}
	temp := uint32((uint32(read[2]) << 16) + (uint32(read[1]) << 8) + uint32(read[0]))
	// change the one bit requested
	if pinVal {
		// set the bit
		temp |= pinMask
	} else {
		// clear the bit
		temp &= ^(pinMask)
	}
	write = make([]byte, 4)
	write[0] = outPortAddr
	write[1] = uint8(temp)
	write[2] = uint8(temp >> 8)
	write[3] = uint8(temp >> 16)
	read = make([]byte, 0)
	if err := ioexp.Tx(write, read); err != nil {
		return errors.New("no reply on I2C buss")
	}
	return nil
}
