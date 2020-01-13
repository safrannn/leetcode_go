package problem1249

import (
	"reflect"
	"testing"
)

func TestProblem1249(t *testing.T) {
	answer :=  "lee(t(c)o)de"
	result := minRemoveToMakeValid("lee(t(c)o)de)")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer =  "ab(c)d"
	result = minRemoveToMakeValid("a)b(c)d")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer =  ""
	result = minRemoveToMakeValid("))((")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = "a(b(c)d)"
	result = minRemoveToMakeValid("(a(b(c)d)")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}