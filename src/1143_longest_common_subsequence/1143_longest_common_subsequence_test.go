package problem1143

import (
	"reflect"
	"testing"
)

func TestProblem1143(t *testing.T) {
	answer := 3
	result := longestCommonSubsequence("abcde", "ace")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 3
	result = longestCommonSubsequence( "abc", "abc")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 0
	result = longestCommonSubsequence("abc","def")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}