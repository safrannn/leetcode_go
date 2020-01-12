package problem1031

import (
	"reflect"
	"testing"
)

func TestProblem1207(t *testing.T) {
	answer := 20
	result := maxSumTwoNoOverlap([]int{0,6,5,2,2,5,1,9,4},1,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 29
	result = maxSumTwoNoOverlap([]int{3,8,1,3,2,1,8,9,0},3,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 31
	result = maxSumTwoNoOverlap([]int{2,1,5,6,0,9,5,0,3,8},4,3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}