package problem0162

import (
	"reflect"
	"testing"
)

// Iterative
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func TestDivisorGame(t *testing.T) {
	a1 := []int{1, 2, 1, 3, 5, 6, 4}
	a1answer := 5
	result1 := findPeakElement(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}
}
