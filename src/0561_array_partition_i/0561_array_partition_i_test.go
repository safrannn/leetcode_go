package problem0561

import (
	"reflect"
	"testing"
)

func TestProblem0561(t *testing.T) {
	a := []int{1432}
	answer := 4
	result := arrayPairSum(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
