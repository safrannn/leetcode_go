package problem0009

import (
	"reflect"
	"testing"
)

func TestProblem0009(t *testing.T) {
	if !reflect.DeepEqual(isPalindrome(121), true) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(isPalindrome(-121), false) {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(isPalindrome(10), false) {
		t.Errorf("wrong") // to indicate test failed
	}
}