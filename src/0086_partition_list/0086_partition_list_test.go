package problem0086


import (
	"reflect"
    "testing"
    "leetcode_go/types"
)

func TestProblem0086(t *testing.T) {
	answer := types.CreateListNode([]int{1,2,2,4,3,5})
	result := partition(types.CreateListNode([]int{1,4,3,2,5,2}),3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}