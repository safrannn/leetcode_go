package problem1133

import (
	"reflect"
	"testing"
)

func TestProblem1133(t *testing.T) {
	answer := 8
	result := largestUniqueNumber([]int{5,7,3,9,4,9,8,3,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = -1
	result = largestUniqueNumber([]int{9,9,8,8})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}