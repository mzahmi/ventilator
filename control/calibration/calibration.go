package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/mzahmi/ventilator/control/adc"
	"github.com/mzahmi/ventilator/control/sensors"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func calibrate() {
	f, err := os.Create("data.txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0}
	for start := time.Now(); time.Since(start) < (time.Second * 30); {
		AdcSlice, err := adc.ReadADC(PIns.AdcID)
		check(err)
		VoltageSignal := AdcSlice[PIns.ID]
		PressureReading := VoltageSignal * 3
		_, err = fmt.Fprintf(w, "Voltage: %v\t Pressure: %v\n", VoltageSignal, PressureReading)
		check(err)
	}
}

func caltest() {
	f, err := os.Create("data.txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = fmt.Fprintf(w, "Voltage, Pressure\n")
	check(err)

	counter := 0
	VoltageSignal := 0
	PressureReading := 0
	for i := 0; i < 10; i++ {
		VoltageSignal = counter + counter
		PressureReading = counter + counter
		_, err = fmt.Fprintf(w, "%v, %v\n", VoltageSignal, PressureReading)
		check(err)
		counter++
		time.Sleep(time.Millisecond * 500)
	}

}

func main() {
	caltest()

}
