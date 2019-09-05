package problem1099

import (
	"reflect"
	"testing"
)

func TestProblem1099(t *testing.T) {
	a := []int{34, 23, 1, 24, 75, 33, 54, 8}
	answer := 58
	result := twoSumLessThanK(a, 60)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
