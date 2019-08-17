package problem0001

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	a := []int{2, 7, 11, 15}
	result := []int{0, 1}
	answer := twoSum(a, 9)
	for i := range answer {
		if answer[i] != result[i] {
			t.Errorf("wrong")
			break
		}
	}
	fmt.Println("the end")
}
