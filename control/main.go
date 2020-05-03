package main

import (
	"fmt"

	"github.com/mzahmi/ventilator/control/controlsystem"
	//"github.com/mzahmi/ventilator/control/sensors"
)

func main() {
	controlsystem.Control()
	fmt.Println("im in main")
}
