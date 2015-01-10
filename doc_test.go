package gps

import (
	"fmt"
)

func ExampleGet() {
	// Data recieved from GPS.
	gprmcData := " $GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*77"

	// Split data on "," to seperate alle values.
	gprmcSplit := strings.Split(gprmcData, "'")

	// gps.Get takes []string generated from stings.Split.
	readableData, _ := Get(gprmcSplit)

	fmt.Printf("%+v\n", readableData)
	// Output: Location object with parsed data.

}

func Example() {
	fmt.Printf("%v", "Test!")
	// Output: Test!

}
