package problem0167


import (
	"reflect"
	"testing"
)

func TestProblem1207(t *testing.T) {
	answer := []int{1,2}
	result := twoSum([]int{2,7,11,15},9)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}