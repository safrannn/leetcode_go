package problem1290

import (
	"leetcode_go/types"
	"reflect"
	"testing"
)

func TestProblem1290(t *testing.T) {
	if !reflect.DeepEqual(getDecimalValue(types.CreateListNode([]int{1,0,1})), 5) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(getDecimalValue(types.CreateListNode([]int{0})), 0) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(getDecimalValue(types.CreateListNode([]int{1})), 1) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(getDecimalValue(types.CreateListNode([]int{1,0,0,1,0,0,1,1,1,0,0,0,0,0,0})), 18880) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(getDecimalValue(types.CreateListNode([]int{0,0})), 0) {
		t.Errorf("wrong") // to indicate test failed
	}
}