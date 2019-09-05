package problem0142

import (
	"leetcode_go/types"
	"reflect"
	"testing"
)


func TestProblem0142(t *testing.T) {
	a := types.CreateListNode([]int{3, 2, 0, -4})
	answer := types.CreateListNode([]int{0,-4})
	result := detectCycle(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
