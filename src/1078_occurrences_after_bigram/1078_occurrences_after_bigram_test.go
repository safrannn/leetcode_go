package problem1078

import (
	"reflect"
	"testing"
)

func TestProblem1078(t *testing.T) {
	answer := []string{"girl","student"}
	result := findOcurrences( "alice is a good girl she is a good student","a","good")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = []string{"we","rock"}
	result = findOcurrences( "we will we will rock you","we","will")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}