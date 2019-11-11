package problem0647

import (
	"reflect"
	"testing"
)

func TestProblem0647(t *testing.T) {
	answer := 3
	result := countSubstrings("abc")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 6
	result = countSubstrings("aaa")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
	
	answer = 9
	result = countSubstrings("abccba")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}