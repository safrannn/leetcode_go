package problem1202

import (
	"reflect"
	"testing"
)

func TestProblem0455(t *testing.T) {
	answer := 6
	result := numWays(3,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}