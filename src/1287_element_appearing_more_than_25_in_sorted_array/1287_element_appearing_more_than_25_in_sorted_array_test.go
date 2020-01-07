package problem1287

import (
	"reflect"
	"testing"
)

func TestProblem1287(t *testing.T) {
	answer := 6
	result := findSpecialInteger([]int{1,2,2,6,6,6,6,7,10})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}
	