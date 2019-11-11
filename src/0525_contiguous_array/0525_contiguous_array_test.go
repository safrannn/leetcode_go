package problem0525

import (
	"reflect"
	"testing"
)

func TestProblem0525(t *testing.T) {
	answer := 2
	result := findMaxLength([]int{0,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 2
	result = findMaxLength([]int{0,1,0})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
	
    answer = 14
	result = findMaxLength([]int{1,0,1,0,1,1,0,1,0,1,0,0,0,1,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}