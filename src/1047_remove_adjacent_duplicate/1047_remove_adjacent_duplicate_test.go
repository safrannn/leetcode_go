package problem1047

import (
	"reflect"
	"testing"
)

func TestProblem1047(t *testing.T) {
	a := "abbaca"
	answer := "ca"
	result := removeDuplicates(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
