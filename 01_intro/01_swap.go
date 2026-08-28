package main

func swap(a int, b int) []int {
	c := a
	a = b
	b = c
	return []int{a, b}
}
