package problem1052

import (
	"reflect"
	"testing"
)

func TestProblem1052(t *testing.T) {
	answer :=  16
	result := maxSatisfied([]int{1,0,1,2,1,1,7,5},[]int{0,1,0,1,0,1,0,1},3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}