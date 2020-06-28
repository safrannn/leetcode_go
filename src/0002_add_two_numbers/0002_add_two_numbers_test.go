package problem0002

import (
	"leetcode_go/types"
	"reflect"
	"testing"
)

func TestProblem0002(t *testing.T) {
	if !reflect.DeepEqual(addTwoNumbers(types.CreateListNode([]int{2,4,3}), types.CreateListNode([]int{5,6,4})), types.CreateListNode([]int{7,0,8})) {
		t.Errorf("wrong") // to indicate test failed
	}
}