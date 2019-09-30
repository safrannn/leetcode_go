package problem1033

import (
	"reflect"
	"testing"
)

func TestProblem1033(t *testing.T) {
	answer := []int{1,2}
	result := numMovesStones(1,2,5)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = []int{0,0}
	result = numMovesStones(4,3,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = []int{1,2}
	result = numMovesStones(3,5,1)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}