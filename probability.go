package main

import "math"

func Exponential(x, l float64) float64 {
	return l * math.Exp(-l*x)
}
