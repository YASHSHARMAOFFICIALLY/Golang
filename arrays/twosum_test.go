package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"pair at start", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"pair at end", []int{3, 2, 4}, 6, []int{1, 2}},
		{"duplicates", []int{3, 3}, 6, []int{0, 1}},
		{"no pair", []int{1, 2, 3}, 100, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
