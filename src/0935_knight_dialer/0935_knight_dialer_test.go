package problem0935

import (
	"reflect"
	"testing"
)

func TestProblem0935(t *testing.T) {
	answer := 10
	result := knightDialer(1)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = 240
	result = knightDialer(5)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = 6576
	result = knightDialer(9)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}


