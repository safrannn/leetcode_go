package problem0539

import (
	"reflect"
	"testing"
)

func TestProblem0539(t *testing.T) {
	answer := 1
	result := findMinDifference([]string{"23:59","00:00"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}