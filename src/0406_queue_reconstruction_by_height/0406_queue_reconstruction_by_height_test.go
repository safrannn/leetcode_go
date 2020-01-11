package problem0406

import (
	"reflect"
	"testing"
)

func TestProblem0406(t *testing.T) {
    answer := [][]int{[]int{5,0},[]int{7,0},[]int{5,2},[]int{6,1},[]int{4,4},[]int{7,1}}
	result := reconstructQueue([][]int{[]int{7,0},[]int{4,4},[]int{7,1},[]int{5,0},[]int{6,1},[]int{5,2}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}