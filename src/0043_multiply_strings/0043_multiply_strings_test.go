package problem0043


import (
	"reflect"
	"testing"
)

func TestProblem0455(t *testing.T) {
	answer := "56088"
	result := multiply("123","456")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}