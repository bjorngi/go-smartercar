package gps

import (
	"fmt"
	"strings"
)

func ExampleGet() {
	// Data recieved from GPS.
	gprmcData := " $GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*77"

	// Split data on "," to seperate alle values.
	gprmcSplit := strings.Split(gprmcData, ",")

	// gps.Get takes []string generated from stings.Split.
	readableData, _ := Get(gprmcSplit)

	// readableData is pointer to Location object
	// &{Speed:1.0443223 Fix:true Coords:{Lon:-74.00695 Lat:40.710236} Bearing:221.11}

	fmt.Printf("%+v\n", readableData)
}

func ExampleChecksum_Valid() {
	// Data recieved from GPS.
	gprmcData := "$GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*77"

	// Test checksum
	fmt.Println(Checksum(gprmcData))
	// Output: true
}

func ExampleChecksum_Invalid() {
	// Data recieved from GPS.
	gprmcData := "$GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*76"

	// Test checksum
	fmt.Println(Checksum(gprmcData))
	// Output: false
}
