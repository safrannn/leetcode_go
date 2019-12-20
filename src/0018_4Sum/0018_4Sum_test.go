package problem0018

import (
	"reflect"
	"testing"
)

func TestProblem0018(t *testing.T) {
	answer := [][]int{[]int{-2, -1, 1, 2},[]int{-2,0,0,2},[]int{-1,  0, 0, 1}}
	result := fourSum([]int{1, 0, -1, 0, -2, 2},0)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}