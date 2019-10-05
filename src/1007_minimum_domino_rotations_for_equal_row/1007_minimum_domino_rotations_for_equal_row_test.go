problem package1007

func minDominoRotations

import (
	"reflect"
	"testing"
)

func TestProblem1007(t *testing.T) {
	answer := 2
	result := minDominoRotations([]int{2,1,2,4,2,2},[]int{5,2,6,2,3,2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = -1
	result = minDominoRotations([]int{3,5,1,2,3},[]int{3,6,3,3,4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}