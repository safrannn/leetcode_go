package problem0274

import (
	"reflect"
	"testing"
)

func TestProblem0274(t *testing.T) {
	answer := 3
	result := hIndex([]int{3,0,6,1,5})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}