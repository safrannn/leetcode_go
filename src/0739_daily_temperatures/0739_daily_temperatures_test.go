package problem0739

import (
	"reflect"
	"testing"
)

func TestProblem0739(t *testing.T) {
	answer := []int{1,1,4,2,1,1,0,0}
	result := dailyTemperatures([]int{73,74,75,71,69,72,76,73})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}