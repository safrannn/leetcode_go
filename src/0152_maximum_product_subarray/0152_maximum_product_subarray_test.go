package problem0152

import (
	"reflect"
	"testing"
)

func TestProblem0152(t *testing.T) {
	a1 := []int{2, 3, -2, 4}
	a1answer := 6
	result1 := maxProduct(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}

	a2 := []int{-2, 0, -1}
	a2answer := 0
	result2 := maxProduct(a2)
	if !reflect.DeepEqual(a2answer, result2) {
		t.Errorf("wrong") // to indicate test failed
	}
}
