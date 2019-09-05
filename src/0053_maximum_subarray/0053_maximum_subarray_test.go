package problem0053

import (
	"reflect"
	"testing"
)

func TestProblem0645(t *testing.T) {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	answer := 6
	result := maxSubArray(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
