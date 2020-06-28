package problem0008

import (
	"reflect"
	"testing"
)

func TestProblem0008(t *testing.T) {
	if !reflect.DeepEqual(myAtoi("42"), 42) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(myAtoi("  -42"), -42) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(myAtoi("4193 with words"), 4193) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(myAtoi("words and 987"), 0) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(myAtoi("-91283472332"), -2147483648) {
		t.Errorf("wrong") // to indicate test failed
	}
}