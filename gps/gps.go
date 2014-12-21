package gps

import (
	"strconv"
	"strings"
)

type Coodrinates struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Location struct {
	Speed   float32     `json:"speed"`
	Fix     bool        `json:"fix"`
	Coords  Coodrinates `json:"coordinates"`
	Bearing float32     `json:"bearing"`
}

func Get(rawLoc string) *Location {
	payload := parse(rawLoc)

	return payload
}

// Parses GPRMC data into Location struct.
// Returns Location struct with JSON converion.
func parse(rawLoc string) *Location {
	locArr := strings.Split(rawLoc, ",")
	loc := buildLocation(locArr)
	return loc
}

// Parses and builds Location struct for GPRMC data.
// Returs Location pointer.
func buildLocation(locArr []string) *Location {
	loc := &Location{
		Fix:     checkfix(locArr[2]),
		Speed:   convertSpeed(locArr[7]),
		Bearing: convertBearing(locArr[8]),
		Coords: Coodrinates{
			Lon: convertDeg(locArr[4], "0"+locArr[3]),
			Lat: convertDeg(locArr[6], locArr[5]),
		},
	}
	return loc
}

func convertBearing(rawBearing string) float32 {
	bearing, err := strconv.ParseFloat(rawBearing, 32)
	if err != nil {
		panic(err)
	}
	return float32(bearing)
}

// checkfix returns if the GPS position is fixed or not as a boolean.
// Invalid input will make it panic.
// Returns boolean, true for fix and false for no fix.
func checkfix(fix string) bool {
	switch fix {
	case "A":
		return true
	case "V":
		return false
	default:
		panic(fix)
	}
}

// Converts from degree-minutes to decimal degree
func convertDeg(heading, rawLon string) float32 {

	deg, _ := strconv.ParseFloat(string(rawLon[0:3]), 32)
	min, _ := strconv.ParseFloat(string(rawLon[3:]), 32)

	decimal := float32(deg) + (float32(min) / 60)
	if heading == "S" || heading == "W" {
		decimal = -decimal
	}

	return decimal
}

// convertSpeed convert knot to m/s.
// Returns float32 of m/s speed
func convertSpeed(knot string) float32 {
	const convertion = float32(0.514444444)
	knotFloat, err := strconv.ParseFloat(knot, 32)
	if err != nil {
		return float32(-1)
	}

	ms := float32(knotFloat) * convertion

	return ms
}
