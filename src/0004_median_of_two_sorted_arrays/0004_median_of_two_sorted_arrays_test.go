package problem0004

import (
	"reflect"
	"testing"
)

func TestProblem0004(t *testing.T) {
	answer := float64(2.0)
	result := findMedianSortedArrays([]int{1,3},[]int{2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = float64(2.5)
	result = findMedianSortedArrays([]int{1,3},[]int{2,4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}


