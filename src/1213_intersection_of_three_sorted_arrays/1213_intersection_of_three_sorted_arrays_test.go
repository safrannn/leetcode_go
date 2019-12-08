package problem1213

import (
	"reflect"
	"testing"
)

func TestProblem1213(t *testing.T) {
	answer := []int{1,5}
	result := arraysIntersection([]int{1,2,3,4,5},[]int{1,2,5,7,9},[]int{1,3,4,5,8})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}