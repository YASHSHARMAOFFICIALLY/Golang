package main

import "math"

func countDigit(n int) int {
	if n == 0 {
		return 1
	}
	value := n
	if value < 0 {
		value = -value
	}
	return int(math.Floor(math.Log10(float64(value)))) + 1
}
