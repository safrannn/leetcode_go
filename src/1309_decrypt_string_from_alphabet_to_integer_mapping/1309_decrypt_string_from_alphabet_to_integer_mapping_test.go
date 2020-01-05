package problem1309

import (
	"reflect"
	"testing"
)

func TestProblem1165(t *testing.T) {
	answer := "jkab"
	result := freqAlphabets("10#11#12")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}
	