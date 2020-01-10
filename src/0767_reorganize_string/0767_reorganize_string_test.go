package problem0767

import (
	"reflect"
	"testing"
)

func TestProblem0739(t *testing.T) {
	answer := "aba"
	result := reorganizeString("aab")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
	
	answer = ""
	result = reorganizeString("aaab")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}