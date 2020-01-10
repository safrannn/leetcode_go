package problem0021

import (
	"leetcode_go/types"
	"reflect"
	"testing"
)


func TestProblem0021(t *testing.T) {
    a := types.CreateListNode([]int{1,2,4})
    b := types.CreateListNode([]int{1,3,4})
	answer := types.CreateListNode([]int{1,1,2,3,4,4})
	result := mergeTwoLists(a,b)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
