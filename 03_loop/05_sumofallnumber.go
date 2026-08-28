package main

func sumofDivisor(n int) int {
	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			pair := n / i
			if pair != i {
				sum += pair
			}
		}
	}
	return sum
}
