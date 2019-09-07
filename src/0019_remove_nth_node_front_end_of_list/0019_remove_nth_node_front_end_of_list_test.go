package problem0019

import (
	"reflect"
	"testing"
	"leetcode_go/types"
)

func TestProblem0019(t *testing.T) {
	a := types.CreateListNode([]int{1, 2, 3, 4, 5})
	answer := types.CreateListNode([]int{1, 2, 3, 5})
	result := removeNthFromEnd(a, 2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}