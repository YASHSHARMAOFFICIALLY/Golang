package main

func isValidTriangle(a float64, b float64, c float64) bool {

	return a+b > c && a+c > b && b+c > a
}
