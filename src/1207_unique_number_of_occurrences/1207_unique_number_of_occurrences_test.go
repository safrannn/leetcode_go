package problem1207

import (
	"reflect"
	"testing"
)

func TestProblem1207(t *testing.T) {
	answer := true
	result := uniqueOccurrences([]int{1,2,2,1,1,3})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = false
	result = uniqueOccurrences([]int{1,2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = true
	result = uniqueOccurrences([]int{-3,0,1,-3,1,1,1,-3,10,0})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}