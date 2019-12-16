package problem1202

import (
	"reflect"
	"testing"
)

func TestProblem0455(t *testing.T) {
	answer := "abcd"
	result := smallestStringWithSwaps("dcab",[][]int{[]int{0,3},[]int{1,2},[]int{0,2}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}