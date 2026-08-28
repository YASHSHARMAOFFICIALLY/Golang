package main

func multiplicationTable(n int) []int {
	result := []int{}
	for i := 1; i <= 10; i++ {
		result = append(result, n*i)
	}
	return result
}
