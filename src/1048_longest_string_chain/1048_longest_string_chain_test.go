package problem1048

import (
	"reflect"
	"testing"
)

func TestProblem1048(t *testing.T) {
	answer := 4
	result := longestStrChain([]string{"a","b","ba","bca","bda","bdca"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}