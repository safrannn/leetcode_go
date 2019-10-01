package problem1180

import (
	"reflect"
	"testing"
)

func TestProblem1180(t *testing.T) {
	answer := 8
	result := countLetters("aaaba")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 55
	result = countLetters("aaaaaaaaaa")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}