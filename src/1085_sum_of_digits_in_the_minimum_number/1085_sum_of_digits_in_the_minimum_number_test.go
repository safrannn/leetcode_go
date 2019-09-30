package problem1085

import (
	"reflect"
	"testing"
)

func TestProblem1085(t *testing.T) {
	answer := 0
	result := sumOfDigits([]int{34,23,1,24,75,33,54,8})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}