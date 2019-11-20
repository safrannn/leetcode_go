package problem0755

import (
	"reflect"
	"testing"
)

func TestProblem0755(t *testing.T) {
	answer := []int{2,2,2,3,2,2,2}
	result := pourWater([]int{2,1,1,2,1,2,2},4,3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = []int{2,3,3,4}
	result = pourWater([]int{1,2,3,4},2,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}