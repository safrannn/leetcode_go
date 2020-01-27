package problem0040

import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	answer :=  [][]int{[]int{1,2,2},[]int{5}}
	result := combinationSum2([]int{2,5,2,1,2},5)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}