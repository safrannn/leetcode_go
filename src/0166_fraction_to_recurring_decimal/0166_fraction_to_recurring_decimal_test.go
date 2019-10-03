package problem0166

import (
	"reflect"
	"testing"
)

func TestProblem0166(t *testing.T) {
	answer := "9.(461538)"
	result := fractionToDecimal(123,13)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = "0.5"
	result = fractionToDecimal(1,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = "0.(6)"
	result = fractionToDecimal(2,3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}