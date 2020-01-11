package problem0991


import (
	"reflect"
	"testing"
)

func TestProblem0091(t *testing.T) {
	answer := 2
	result := brokenCalc(2,3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 2
	result = brokenCalc(5,8)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 3
	result = brokenCalc(3,10)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    answer = 1023
	result = brokenCalc(1024,1)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}