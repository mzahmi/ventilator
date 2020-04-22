package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

// Reading ...  reads from usb
// must chmod /dev/ttyACM0
func Reading() string {
	textSend := ""
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 10)
	n, err := s.Read(buf)
	sensVal := string(buf[:n])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%q", buf[:n])
	fmt.Println(sensVal)
	textSend = sensVal
	return textSend

}
