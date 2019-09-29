package problem1200

import (
	"reflect"
	"testing"
)

func TestProblem1200(t *testing.T) {
	answer := [][]int{[]int{1,3}}
	result := minimumAbsDifference([]int{1,3,6,10,15})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}