package main

func classify(n int) string {
	if n > 0 {
		return "positive"
	} else if n == 0 {
		return "zero"
	} else {
		return "negative"
	}
}
