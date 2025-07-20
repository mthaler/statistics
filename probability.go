package main

import "math"

/*
Assumes a constant decay rate, an example is the time until a radioactive particle decays, or the time between clicks of a Geiger counter
*/
func Exponential(x, l float64) float64 {
	return l * math.Exp(-l*x)
}
