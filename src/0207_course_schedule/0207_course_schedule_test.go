package problem0207

import (
	"reflect"
	"testing"
)

func TestProblem0207(t *testing.T) {
	answer := true
	result := canFinish(2,[][]int{[]int{1,0}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

	answer = false
	result = canFinish(2,[][]int{[]int{1,0},[]int{0,1}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}