package problem0034

import (
	"reflect"
	"testing"
)

func TestProblem0018(t *testing.T) {
	answer := []int{3,4}
	result := searchRange([]int{5,7,7,8,8,10},8)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}