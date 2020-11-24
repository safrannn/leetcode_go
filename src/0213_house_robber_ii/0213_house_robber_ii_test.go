package problem0213

import (
	"reflect"
	"testing"
)

func TestProblem0123(t *testing.T) {
	answer := 3
	result := rob([]int{2,3,2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 4
	result = rob([]int{1,2,3,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 0
	result = uniqueOccurrences([]int{0})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}