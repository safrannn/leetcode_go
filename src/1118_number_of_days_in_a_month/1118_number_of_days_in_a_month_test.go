package problem1118

import (
	"reflect"
	"testing"
)

func TestProblem1118(t *testing.T) {
	answer := 31
	result := numberOfDays(1992,7)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    answer = 29
	result = numberOfDays(2000,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    answer = 28
	result = numberOfDays(1900,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}