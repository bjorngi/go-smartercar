// Package gps provides parsing of $GPRMC GPS-data and NMEA checksum checking.
// Only retuns data on fixed GPRMC data.
package gps

import (
	"errors"
	"strconv"
)

type Coordinates struct {
	Lon float32 `json:"lon"` // Longtitude in decimal format.
	Lat float32 `json:"lat"` // Latitude in decial format.
}

type Location struct {
	Speed   float32     `json:"speed"` // Speed in m/s.
	Fix     bool        `json:"fix"`   // Boolean to verify gps-fix.
	Coords  Coordinates `json:"coordinates"`
	Bearing float32     `json:"bearing"` // Bearing in degrees.
}

// Get takes in $GPRMC data split up into string array.
// Parses data and converting to decimal degrees.
// Retuns Location object with all parsed infomation.
func Get(locArr []string) (*Location, error) {
	if checkfix(locArr[2]) {
		payload := parse(locArr)
		return payload, nil
	}
	err := errors.New("GPS data not fixed")
	return nil, err

}

// Parses GPRMC data into Location struct.
// Returns Location struct with JSON converion.
func parse(locArr []string) *Location {
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
		Coords: Coordinates{
			Lat: convertDeg(locArr[4], "0"+locArr[3]),
			Lon: convertDeg(locArr[6], locArr[5]),
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
