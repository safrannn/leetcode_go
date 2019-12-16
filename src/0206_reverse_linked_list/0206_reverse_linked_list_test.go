package problem0206

import (
	"leetcode_go/types"
	"reflect"
	"testing"
)


func TestProblem0206(t *testing.T) {
	answer := types.CreateListNode([]int{5,4,3,2,1})
	result := reverseList(types.CreateListNode([]int{1,2,3,4,5}))
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
