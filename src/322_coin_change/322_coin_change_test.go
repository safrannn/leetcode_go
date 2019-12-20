package problem0322


import (
	"reflect"
	"testing"
)

func TestProblem0322(t *testing.T) {
	answer := 3
	result := coinChange([]int{1, 2, 5},11)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = -1
	result = coinChange([]int{2},3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}