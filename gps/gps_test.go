package gps

import (
	"fmt"
	"testing"
)

var GPRMC string = " $GPRMC,194509.000,A,4042.6142,N,07400.4168,W,2.03,221.11,160412,,,A*77"

func TestLocationParsing(t *testing.T) {
	loc := parse(GPRMC)
	fmt.Printf("Values recieved: %#v", loc)

	if !loc.Fix {
		t.Errorf("Location not fixed, returned: %v", loc.Fix)
	}
}

func TestCheckFix(t *testing.T) {
	a := checkfix("A")
	if !a {
		t.Errorf("Checkfix failed: returned %#v", a)
	}

	b := checkfix("V")
	if b {
		t.Errorf("Checkfix failed: returned %#v", b)
	}
}

func TestConvertSpeed(t *testing.T) {
	knot := float32(35)
	assMS := float32(18.0055556)
	ms := convertSpeed("35")

	if ms != assMS {
		t.Errorf("ConvertSpeed failed,%v knot should return %v, retured %v", knot, assMS, ms)
	}
}

// Benchmarking
func BenchmarkParsing(*testing.B) {
	parse(GPRMC)
}
