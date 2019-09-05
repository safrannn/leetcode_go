package problem0994

import (
	"reflect"
	"testing"
)

func TestProblem0994(t *testing.T) {
	a := [][]int{[]int{2, 1, 1}, []int{0, 1, 1}, []int{1, 0, 1}}
	answer := -1
	result := orangesRotting(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
