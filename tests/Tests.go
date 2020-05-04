package tests

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/mzahmi/ventilator/control/sensors"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// PressureTest tests the reading of the pressure sensors for 1 minute
// on a 1 second intervals
func PressureTest(PS *sensors.Pressure) {
	f, err := os.Create("PressureTest.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(60 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Sensor calibration done")
			return
		case t := <-ticker.C:
			_, err = fmt.Fprintf(w, "%v\tThe pressure reading from %s is %v\n", t, PS.Name, PS.ReadPressure())
			check(err)
		}
	}
}

// FlowTest tests the reading of the flow sensors for 1 minute
// on a 1 second intervals
func FlowTest(FS *sensors.Flow) {
	f, err := os.Create("FlowTest.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(60 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Sensor calibration done")
			return
		case t := <-ticker.C:
			_, err = fmt.Fprintf(w, "%v\tThe pressure reading from %s is %v\n", t, FS.Name, FS.ReadFlow())
			check(err)
		}
	}
}
