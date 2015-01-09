// Package serial provides connection to arduino serial.
package serial

import (
	"bufio"
	"github.com/huin/goserial"
	"log"
)

// ReadArduino takes in a channel
// Opens serial comminucation with arduino and reads data.
// Retunrs data over channel as string.
func ReadArduino(ch chan string) {
	c := &goserial.Config{Name: "/dev/ttyACM3", Baud: 9600}

	s, err := goserial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		text := scanner.Text()

		ch <- text
	}
}
