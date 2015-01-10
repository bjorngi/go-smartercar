package gps

import (
	"fmt"
	"testing"
)

// Example
func ExampleExamples() {
	gprmcData := " $GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*77"
	gprmcSplit := strings.Split(gprmcData, "'")
	readableData, _ := Get(gprmcSplit)
	fmt.Printf("%+v\n", readableData)
	//Output: Location object with parsed data.

}
