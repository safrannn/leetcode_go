package problem0287

import (
	"reflect"
	"testing"
)

func TestProblem0287(t *testing.T) {
	a1 := []int{1, 3, 4, 2, 2}
	a1answer := 2
	result1 := findDuplicate(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}

	a2 := []int{3, 1, 3, 4, 2}
	a2answer := 3
	result2 := findDuplicate(a2)
	if !reflect.DeepEqual(a2answer, result2) {
		t.Errorf("wrong") // to indicate test failed
	}
}
