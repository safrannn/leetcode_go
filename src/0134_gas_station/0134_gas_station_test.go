package problem0134


import (
	"reflect"
	"testing"
)

func TestProblem0134(t *testing.T) {
	answer := 3
	result := canCompleteCircuit([]int{1,2,3,4,5},[]int{3,4,5,1,2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = -1
	result = canCompleteCircuit([]int{2,3,4},[]int{3,4,3})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}