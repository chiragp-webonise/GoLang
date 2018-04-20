// Go structure usage example based on Mars
// Planned for Chapter 4 of Learn Go http://yng.mn/learngolang
// Copyright (c) 2015 Nathan Youngman, see https://github.com/nathany/learngo

package main

import (
	"fmt"
	"math"
)

// point with a latitude, longitude (positive East).
type point struct {
	lat  float64
	long float64
}

// radius of Mars in kilometers (volumetric mean).
const radius = 3389.5

// distance calculation using the Spherical Law of Cosines.
func distance(p1, p2 point) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	spirit := point{85.28284799367495,25.08959000813778} // At Columbia Memorial Station
	opportunity := point{85.29403545452726,25.08912193086729} // At Challenger Memorial Station

	fmt.Printf("%.2f km", distance(spirit, opportunity))
}
