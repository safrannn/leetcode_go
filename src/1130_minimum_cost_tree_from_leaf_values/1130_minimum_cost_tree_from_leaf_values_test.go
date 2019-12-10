package problem1130

import (
	"reflect"
	"testing"
)

func TestProblem1130(t *testing.T) {
	answer := 560
	result := mctFromLeafValues([]int{15,13,5,6,2,4,15})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 32
	result = mctFromLeafValues([]int{6,2,4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}