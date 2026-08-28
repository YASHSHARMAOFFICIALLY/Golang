package main

func divide(dividend int, divisor int) []int {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return []int{quotient, remainder}
}
