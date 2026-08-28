package main

func sumByparity(n int, parity string) int {
	result := 0
	if parity == "even" {
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				result += i
			}
		}
	}
	if parity == "odd" {
		for i := 1; i <= n; i++ {
			if i%2 != 0 {
				result += i
			}
		}
	}

	return result
}
