package adc

import (
	"encoding/binary"
	"errors"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

const numChannels = 8
const adc1_cs_GPIO = "GPIO20"
const adc2_cs_GPIO = "GPIO21"
const adc3_cs_GPIO = "GPIO26"

func ReadADC(adcID uint8) ([]float32, error) {

	adcReadResult := make([]float32, numChannels)

	// init the drivers on the Linux host
	if _, err := host.Init(); err != nil {
		return adcReadResult, errors.New("periph.io host.Init() failed")
	}

	var cs_GPIO string
	switch adcID {
	case 1:
		cs_GPIO = adc1_cs_GPIO
	case 2:
		cs_GPIO = adc2_cs_GPIO
	case 3:
		cs_GPIO = adc3_cs_GPIO
	}

	chipSel := gpioreg.ByName(cs_GPIO)
	if chipSel == nil {
		return adcReadResult, errors.New("failed to access chip select GPIO")
	}
	// resting state of chip select is high
	chipSel.Out(gpio.High)

	spiPort, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		return adcReadResult, errors.New("failed to access spi0")
	}

	defer spiPort.Close()

	// Convert the spi.Port into a spi.Conn so it can be used for communication.
	spiConn, err := spiPort.Connect(1*physic.MegaHertz, spi.NoCS, 8)
	
	if err != nil {
		return adcReadResult, errors.New("failed to connect to spi0")
	}

	// first 16 bits starts the conversion, next 16 bit are where results are clocked out,
	// therefore you need to send in 9 words to get 8 words of data out see pages 16 and 17
	// of the data sheet
	write := make([]byte, numChannels*2+2)
	read := make([]byte, numChannels*2+2)

	// set the channel address's (see page 8 of ADC128S102 data sheet for details)
	for ii := 0; ii < numChannels; ii++ {
		write[ii*2] = byte(ii << 3)
	}

	// write and read on SPI
	chipSel.Out(gpio.Low)
	if err := spiConn.Tx(write, read); err != nil {
		return adcReadResult, errors.New("failed to read ADC")
	}
	chipSel.Out(gpio.High)

	for ii := 0; ii < numChannels; ii++ {
		// note that first word is skipped because second word is where valid channel 0 data is
		adcUint := binary.BigEndian.Uint16(read[2*ii+2 : 2*ii+2+2])
		// 12 bit ADC with 3.3 volt ref voltage
		adcReadResult[ii] = float32(adcUint) * float32(3.3/4096.0)
	}
	return adcReadResult, nil
}
