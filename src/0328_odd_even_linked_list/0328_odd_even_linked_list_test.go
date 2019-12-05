package problem0328

import (
	"leetcode_go/types"
	"reflect"
	"testing"
)


func TestProblem0328(t *testing.T) {
	a := types.CreateListNode([]int{1,2,3,4,5})
	answer := types.CreateListNode([]int{1,3,5,2,4})
	result := oddEvenList(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    a = types.CreateListNode([]int{2,1,3,5,6,4,7})
	answer = types.CreateListNode([]int{2,3,6,7,1,5,4})
	result = oddEvenList(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}
