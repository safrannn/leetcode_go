package problem0547

import (
	"reflect"
	"testing"
)

func TestProblem0547(t *testing.T) {
	a := [][]int{[]int{1,1,0},[]int{1,1,0},[]int{0,0,1},}
	answer := 2
	result := findCircleNum(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
