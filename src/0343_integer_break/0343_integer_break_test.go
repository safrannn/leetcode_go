package problem0343

import (
	"reflect"
	"testing"
)

func TestProblem1249(t *testing.T) {
	answer :=  1
	result := integerBreak(2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer =  36
	result = integerBreak(10)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}