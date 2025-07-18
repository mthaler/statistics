package main

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
	l := len(n)
	return n[l/2]
}
