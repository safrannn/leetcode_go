package problem1150

import (
	"reflect"
	"testing"
)

func TestProblem1150(t *testing.T) {
	answer := true
	result := isMajorityElement([]int{2,4,5,5,5,5,5,6,6},5)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = false
	result = isMajorityElement([]int{10,100,101,101},101)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}