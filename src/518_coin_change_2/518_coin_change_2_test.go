package problem0518


import (
	"reflect"
	"testing"
)

func TestProblem0518(t *testing.T) {
	answer := 4
	result := change(5,[]int{1,2,5})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 0
	result = change(3,[]int{2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 1
	result = change(10,[]int{10})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}