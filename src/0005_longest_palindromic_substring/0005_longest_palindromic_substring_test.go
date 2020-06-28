package problem0005

import (
	"reflect"
	"testing"
)

func TestProblem0005(t *testing.T) {
	if !reflect.DeepEqual(longestPalindrome("babad"), "bab") {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(longestPalindrome("cbbd"), "bb") {
		t.Errorf("wrong") // to indicate test failed
	}
}