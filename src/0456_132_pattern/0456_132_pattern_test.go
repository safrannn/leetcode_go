package problem0456

import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	answer :=  false
	result := find132pattern([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer =  true
	result = find132pattern([]int{3, 1, 4, 2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer =  true
	result = find132pattern([]int{-1, 3, 2, 0})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}