package problem0216

import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	answer :=  [][]int{[]int{1,2,4}}
	result := combinationSum3(3,7)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer =  [][]int{[]int{1,2,6},[]int{1,3,5},[]int{2,3,4}}
	result = combinationSum3(3,9)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}