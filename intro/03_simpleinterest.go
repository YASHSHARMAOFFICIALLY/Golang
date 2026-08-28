package main

import "math"

func interest(principal float64, rate float64, time float64) []float64 {

	simpleInterest := principal * rate * time / 100
	compoundInterest := principal*math.Pow(1+rate/100, time) - principal

	simpleInterest = math.Round(simpleInterest*100) / 100
	compoundInterest = math.Round(compoundInterest*100) / 100

	return []float64{simpleInterest, compoundInterest}

}
