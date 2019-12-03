package problem0033

import (
	"reflect"
	"testing"
)

func TestProblem0033(t *testing.T) {
	answer := 4
	result := search([]int{4,5,6,7,0,1,2},0)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = -1
	result = search([]int{4,5,6,7,0,1,2},3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}

	answer = 0
	result = search([]int{5,1,3},5)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}