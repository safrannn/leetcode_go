package problem1342

import (
	"reflect"
	"testing"
)

func TestProblem1342(t *testing.T) {
	if !reflect.DeepEqual(numberOfSteps(14), 6) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(numberOfSteps(8), 4) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(numberOfSteps(123), 12) {
		t.Errorf("wrong") // to indicate test failed
    }
}