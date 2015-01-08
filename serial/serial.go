// Package main provides connection to arduino serial.
package serial

import (
	"bufio"
	"github.com/huin/goserial"
	"log"
)

func readArduino(ch chan string) {
	c := &goserial.Config{Name: "/dev/ttyACM3", Baud: 9600}

	s, err := goserial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
}
