package dac

import (
	"encoding/binary"
	"errors"
	"math"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

const dac1_cs_GPIO = "GPIO8"
const dac2_cs_GPIO = "GPIO7"
const dacZeroOutWord = 0x7000

// see page 17 for word details

func WriteDac(dacID uint8, dacChan uint8, voltage float64) error {

	if dacID < 1 || dacID > 2 {
		return errors.New("dacID must be 1 or 2")
	}

	// map outA to 0, ... outD to 3
	if dacChan > 3 {
		return errors.New("dacChan must be in range of 0 to 3")
	}

	if voltage < 0.0 || voltage > 10.0 {
		return errors.New("voltage must be between 0 and 10 volts")
	}

	// init the drivers on the Linux host
	if _, err := host.Init(); err != nil {
		return errors.New("periph.io host.Init() failed")
	}

	var cs_GPIO string
	if dacID == 1 {
		cs_GPIO = dac1_cs_GPIO
	} else {
		cs_GPIO = dac2_cs_GPIO
	}

	chipSel := gpioreg.ByName(cs_GPIO)
	if chipSel == nil {
		return errors.New("failed to access chip select GPIO")
	}

	// 8 bit DAC with 3.3 volt reference
	// opamp gain = 1 + 7.15/3.32 = 3.154
	dacWord := uint16(math.Round(((voltage / 3.154) / 3.3) * 255))
	// see data sheet page 17 for bit twiddling details
	dacWord = dacWord << 4
	// set OP1 to 1 and OP0 to zero (Write to specified register and update outputs)
	dacWord |= 0x1000
	// set A1 and A0 bits for DAC channel requested
	switch dacChan {
	case 0:
		dacWord |= 0x0000 //this one is really a nop
	case 1:
		dacWord |= 0x4000
	case 2:
		dacWord |= 0x8000
	case 3:
		dacWord |= 0xC000
	}

	// resting state of chip select is high
	chipSel.Out(gpio.High)

	spiPort, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		return errors.New("failed to access spi0")
	}

	defer spiPort.Close()

	// Convert the spi.Port into a spi.Conn so it can be used for communication.
	spiConn, err := spiPort.Connect(10*physic.MegaHertz, spi.NoCS | spi.Mode1, 8)
	if err != nil {
		return errors.New("failed to connect to spi0")
	}

	write := make([]byte, 2)
	read := make([]byte, 0)

	binary.BigEndian.PutUint16(write, dacWord)

	chipSel.Out(gpio.Low)
	if err := spiConn.Tx(write, read); err != nil {
		return errors.New("failed to write to DAC on SPI")
	}
	chipSel.Out(gpio.High)

	return nil
}

func DacsAllZeroOut() error {
	// init the drivers on the Linux host
	if _, err := host.Init(); err != nil {
		return errors.New("periph.io host.Init() failed")
	}
	chipSel1 := gpioreg.ByName(dac1_cs_GPIO)
	if chipSel1 == nil {
		return errors.New("failed to access chip select 1 GPIO")
	}
	chipSel2 := gpioreg.ByName(dac2_cs_GPIO)
	if chipSel2 == nil {
		return errors.New("failed to access chip select 2 GPIO")
	}

	chipSel1.Out(gpio.High)
	chipSel2.Out(gpio.High)

	spiPort, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		return errors.New("failed to access spi0")
	}

	defer spiPort.Close()

	// Convert the spi.Port into a spi.Conn so it can be used for communication.
	spiConn, err := spiPort.Connect(1*physic.MegaHertz, spi.NoCS, 8)
	if err != nil {
		return errors.New("failed to connect to spi0")
	}

	// this will place an internal 2.5K resistor to ground at all DAC outputs (see page 16 of data sheet)
	// this will in turn drive all the opamp outputs to "close" to zero volts
	write := []byte{0x70, 0x00}
	read := make([]byte, 0)

	// zero out DAC1
	chipSel1.Out(gpio.Low)
	if err := spiConn.Tx(write, read); err != nil {
		return errors.New("failed to write to DAC on SPI")
	}
	chipSel1.Out(gpio.High)

	// zero out DAC2
	chipSel2.Out(gpio.Low)
	if err := spiConn.Tx(write, read); err != nil {
		return errors.New("failed to write to DAC on SPI")
	}
	chipSel2.Out(gpio.High)

	return nil
}
