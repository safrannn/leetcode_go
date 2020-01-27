package problem0377

import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	answer :=  7
	result := combinationSum4([]int{1,2,3},4)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}