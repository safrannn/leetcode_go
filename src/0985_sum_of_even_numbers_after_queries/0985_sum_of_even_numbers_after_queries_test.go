package problem0985

import (
	"reflect"
	"testing"
)

func TestProblem0985(t *testing.T) {
    a := []int{1,2,3,4}
    aQueries := [][]int{[]int{1,0},[]int{-3,1},[]int{-4,0},[]int{2,3}}
	answer := []int{8,6,2,4}
	result := sumEvenAfterQueries(a,aQueries)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}