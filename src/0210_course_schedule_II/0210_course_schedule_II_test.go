package problem0210

import (
	"reflect"
	"testing"
)

func TestProblem0210(t *testing.T) {
	answer := []int{0,1}
	result := findOrder(2,[][]int{[]int{1,0}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = []int{0,1,2,3}
	result = findOrder(4,[][]int{[]int{1,0},[]int{2,0},[]int{3,1},[]int{3,2}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}