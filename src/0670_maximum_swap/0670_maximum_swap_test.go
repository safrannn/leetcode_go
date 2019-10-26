package problem0670


import (
	"reflect"
	"testing"
)

func TestProblem0670(t *testing.T) {
	answer := 9973
	result := maximumSwap(9973)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 7236
	result = maximumSwap(2736)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}