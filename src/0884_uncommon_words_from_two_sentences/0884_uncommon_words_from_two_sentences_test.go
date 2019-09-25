package problem0884

import (
	"reflect"
	"testing"
)

func TestProblem0605(t *testing.T) {
    a1 := "apple apple"
    a2 := "banana"
	answer := []string{"banana"}
	result := uncommonFromSentences(a1, a2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    b1 := "this apple is sweet"
    b2 := "this apple is sour"
	answer = []string{"sweet","sour"}
	result = uncommonFromSentences(b1, b2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
