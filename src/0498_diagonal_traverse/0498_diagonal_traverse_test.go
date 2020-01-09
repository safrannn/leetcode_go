package problem0498


import (
	"reflect"
	"testing"
)

func TestProblem0498(t *testing.T) {
	a1 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	result1 := findDiagonalOrder(a1)
	a1answer := []int{1,2,4,7,5,3,6,8,9}
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}
}
