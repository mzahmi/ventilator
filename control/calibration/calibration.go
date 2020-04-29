package calibration

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mzahmi/ventilator/control/adc"
	"github.com/mzahmi/ventilator/control/sensors"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func calibrate() {
	f, err := os.Create("/data")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	PIns := sensors.Pressure{Name: "SNS_P_INS", ID: 0, AdcID: 1, MMH2O: 0}
	for {
		AdcSlice, err := adc.ReadADC(PIns.AdcID)
		check(err)
		VoltageSignal := AdcSlice[PIns.ID]
		PressureReading := VoltageSignal * 3
		_, err = fmt.Fprintf(w, "Voltage: %v\t Pressure: %v\n", VoltageSignal, PressureReading)
		check(err)
	}

}
