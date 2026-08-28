package main

func grade(marks int) string {

	if marks >= 90 {
		return "A"
	} else if marks >= 80 {
		return "B"
	} else if marks >= 70 {
		return "C"
	} else if marks >= 60 {
		return "D"
	}
	return "F"
}
