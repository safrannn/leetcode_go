package problem0393

import (
	"reflect"
	"testing"
)

func TestProblem1207(t *testing.T) {
    answer := true
	result := validUtf8([]int{197, 130, 1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = false
	result = validUtf8([]int{235, 140, 4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}