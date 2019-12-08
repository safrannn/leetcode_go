package problem1221

import (
	"reflect"
	"testing"
)

func TestProblem1221(t *testing.T) {
	answer := 4
	result := balancedStringSplit("RLRRLLRLRL")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 3
	result = balancedStringSplit("RLLLLRRRLR")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 1
	result = uniqueOccurrences("LLLLRRRR")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}