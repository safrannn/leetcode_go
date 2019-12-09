package problem0039

import (
	"reflect"
	"testing"
)

func TestProblem0039(t *testing.T) {
	answer := [][]int{[]int{3,2,2},[]int{7}}
	result := combinationSum([]int{2,3,6,7},7)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = [][]int{[]int{2,2,2,2}, []int{3,3,2},[]int{5,3}}
	result = combinationSum([]int{2,3,5},8)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}