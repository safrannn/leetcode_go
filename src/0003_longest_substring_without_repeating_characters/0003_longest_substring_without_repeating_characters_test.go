package problem0003


import (
	"reflect"
	"testing"
)

func TestProblem0003(t *testing.T) {
	answer := 3
	result := lengthOfLongestSubstring("abcabcbb")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 1
	result = lengthOfLongestSubstring("bbbb")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 3
	result = lengthOfLongestSubstring("pwwkew")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}