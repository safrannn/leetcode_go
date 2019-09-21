package problem0754

import (
	"reflect"
	"testing"
)

func TestProblem0754(t *testing.T) {
	a := 3
	answer := 2
	result := reachNumber(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}