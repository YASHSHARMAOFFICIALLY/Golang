// Package arrays holds array/slice problems.
package arrays

// TwoSum returns the indices of the two numbers in nums that add up to target,
// or nil if no such pair exists.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
