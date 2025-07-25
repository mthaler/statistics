package main

import (
	"math"
	"sort"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Mean[N Number](n ...N) float64 {
	var sum float64
	for _, v := range n {
		sum += float64(v)
	}
	sum = sum / float64(len(n))
	return sum
}

func Median[N Number](n ...N) N {
	s := n[:]
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	l := len(n)
	return s[l/2]
}

func Variance[N Number](n ...N) float64 {
	var result float64
	m := Mean(n...)
	for _, v := range n {
		result += (float64(v) - m) * (float64(v) - m)
	}
	return result / float64(len(n))
}

func StandardDeviation[N Number](n ...N) float64 {
	return math.Sqrt(Variance(n...))
}

func Skewness[N Number](n ...N) float64 {
	mean := Mean(n...)
	median := float64(Median(n...))
	sd := StandardDeviation(n...)
	return 3 * (mean - median) / sd
}
