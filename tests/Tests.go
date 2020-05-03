package tests

import (
	"bufio"
	"log"
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
	defer f.close()
	w := bufio.NewWriter(f)

	for start := time.Now(); time.Since(start) < (time.Second * 60); {
		for x := range time.Tick(time.Second) {
			_, err = fmt.Fprintf("%t\tThe pressure reading from %s is %v\n", time.Now(), PS.Name, PS.ReadPressure())
			check(err)
		}
	}
}

// FlowTest tests the reading of the flow sensors for 1 minute
// on a 1 second intervals
func FlowTest(FS *sensors.Flow) {
	f, err := os.Create("FlowTest.txt")
	check(err)
	defer f.close()
	w := bufio.NewWriter(f)

	for start := time.Now(); time.Since(start) < (time.Second * 60); {
		for x := range time.Tick(time.Second) {
			_, err = fmt.Fprintf("%t\tThe pressure reading from %s is %v\n", time.Now(), FS.Name, FS.ReadFlow())
			check(err)
		}
	}
}
