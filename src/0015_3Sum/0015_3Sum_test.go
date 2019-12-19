package problem0015


import (
	"reflect"
	"testing"
)

func TestProblem0015(t *testing.T) {
	answer := [][]int{[]int{-2,0,2},[]int{-2,1,1}}
	result := threeSum([]int{-2,0,1,1,2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer := [][]int{[]int{-1,0,1},[]int{-1,-1,2}}
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}