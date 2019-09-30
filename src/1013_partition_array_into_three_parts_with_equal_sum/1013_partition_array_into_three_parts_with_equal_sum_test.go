package problem1013

import (
	"reflect"
	"testing"
)

func TestProblem1013(t *testing.T) {
	answer := true
	result := canThreePartsEqualSum([]int{0,2,1,-6,6,-7,9,1,2,0,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = false
	result = canThreePartsEqualSum([]int{0,2,1,-6,6,7,9,-1,2,0,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = true
	result = canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}