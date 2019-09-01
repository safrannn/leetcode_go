package problem0075

import (
	"reflect"
	"testing"
)

func sortColors(nums []int) []int {
	leftBound, rightBound := 0, len(nums)-1

	for i := 0; i <= rightBound; i++ {
		if nums[i] == 0 {
			nums[i], nums[leftBound] = nums[leftBound], nums[i]
			leftBound++
		} else if nums[i] == 2 {
			nums[i], nums[rightBound] = nums[rightBound], nums[i]
			rightBound--
			i--
		}
	}
	return nums
}

func TestDivisorGame(t *testing.T) {
	a1 := []int{2, 0, 2, 1, 1, 0}
	a1answer := []int{0, 0, 1, 1, 2, 2}
	result1 := sortColors(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}
}
