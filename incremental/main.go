package main

import (
	"math"
)

func Solution(a []int) int {
	count, sum := 0.0, 0.0
	for _, n := range a {
		sum += float64(n)
	}
	avg := math.Round(sum / float64(len(a)))
	for _, n := range a {
		count += math.Abs(float64(n) - avg)
	}
	return int(count)
}
