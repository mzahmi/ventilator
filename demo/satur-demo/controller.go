package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

var (
	//main valve
	MV = gpioreg.ByName("GPIO17")
	//inhale valve
	MIns = gpioreg.ByName("GPIO27")
	//exhale valve
	MExp = gpioreg.ByName("GPIO22")
	// BCT
	BCT float32
	// TI
	TI        float32
	TE        float32
	IR        float32
	ER        float32
	BPM       float32
	StartBool = true
)

func StartClicked() {
	fmt.Println("Start Clicked")
	// MV = gpioreg.ByName("GPIO17")
	MV.Out(gpio.High)
	StartBool = false
}

func StopClicked() {
	fmt.Println("Stop Clicked")
	// MV = gpioreg.ByName("GPIO17")
	MV.Out(gpio.Low)
	StartBool = true
}

func CalculateBCT(bpm float32) {
	BPM = bpm
	BCT = 60 / (BPM)
	IR = 1
	ER = 2
	TI = ((IR * BCT) / (IR + ER))
	TE = BCT - TI

	fmt.Println(BPM, BCT, IR, ER, TI, TE)
}

func InitiateHost() {
	_, err := host.Init()
	if err != nil {
		log.Fatalf("failed to initialize periph: %v", err)
	}
	fmt.Println("Host initiated")
	//main valve
	MV = gpioreg.ByName("GPIO17")
	//inhale valve
	MIns = gpioreg.ByName("GPIO27")
	//exhale valve
	MExp = gpioreg.ByName("GPIO22")

}
