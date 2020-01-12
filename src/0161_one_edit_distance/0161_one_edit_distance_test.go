package problem0161

import (
	"reflect"
	"testing"
)

func TestProblem0161(t *testing.T) {
	answer := true
	result := isOneEditDistance("ab","acb")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = true
	result = isOneEditDistance("ab","cb")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = true
	result = isOneEditDistance("1203","1213")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}