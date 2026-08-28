package main

func convert(temp float64, scale string) float64 {
	if scale == "C" {
		temp = temp*9/5 + 32
	} else {
		temp = (temp - 32) * 5 / 9
	}
	return temp
}
