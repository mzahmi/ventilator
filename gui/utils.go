package main

import (
	"fmt"
	"io/ioutil"
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

// ReadFromFile reads text from file "test.txt"
func ReadFromFile() string {
	returnString := ""

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		returnString = "No file found"
	} else {
		returnString = string(data)
	}
	return returnString
}
