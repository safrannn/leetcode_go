package problem0153

import (
	"reflect"
	"testing"
)

func TestProblem0153(t *testing.T) {
	a1 := []int{3, 4, 5, 1, 2}
	a1answer := 1
	result1 := findMin(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}

	a2 := []int{4, 5, 6, 7, 0, 1, 2}
	a2answer := 0
	result2 := findMin(a2)
	if !reflect.DeepEqual(a2answer, result2) {
		t.Errorf("wrong") // to indicate test failed
	}
}
