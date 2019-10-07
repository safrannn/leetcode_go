package problem0096

import (
	"reflect"
	"testing"
)

func TestProblem0096(t *testing.T) {
	answer := 5
	result := numTrees(3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 208012
	result = numTrees(12)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 1
	result = numTrees(1)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}