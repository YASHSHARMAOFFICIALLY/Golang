package main

func printNumber(n int) []int {
	result := []int{}
	for i := 0; i <= n; i++ {
		result = append(result, i)
	}
	return result
}
