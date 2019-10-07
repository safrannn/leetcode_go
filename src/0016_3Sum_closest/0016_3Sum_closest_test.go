package problem0016

import (
	"reflect"
	"testing"
)

func TestProblem0016(t *testing.T) {
	answer := 2
	result := threeSumClosest([]int{-1,2,1,-4},1)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}