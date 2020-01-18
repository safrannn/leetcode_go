package problem0373

import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	answer :=  [][]int{[]int{1,2},[]int{1,4},[]int{1,6}}
	result := kSmallestPairs([]int{1, 7,11}, []int{2,4,6},3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer =  [][]int{[]int{1,2},[]int{1,4},[]int{1,6}}
	result = kSmallestPairs([]int{1, 7,11}, []int{2,4,6},3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer =  [][]int{[]int{1,3},[]int{2,3}}
	result = kSmallestPairs([]int{1,2}, []int{3},3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}


