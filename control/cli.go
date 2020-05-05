package main

import (
	"fmt"
)

// Implementation of a command line interface
func main() {
	var Option int
	fmt.Println("Welcome to DFF Ventilator CLI")
	fmt.Println(`Please Select Vent Mode from the following:
	1 = Volume AC
	2 = Pressure AC
	3 = PSV
	4 = Exit`)
	switch Option {
	case 1:
		fmt.Println("not available yet")
	case 2:
		fmt.Println("Please input the following:")
	}

}
