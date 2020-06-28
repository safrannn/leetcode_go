package problem0007

import (
	"reflect"
	"testing"
)

func TestProblem0007(t *testing.T) {
	if !reflect.DeepEqual(reverse(123), 321) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(reverse(-123), -321) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(reverse(120), 21) {
		t.Errorf("wrong") // to indicate test failed
	}
}