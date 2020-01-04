package problem1299

import (
	"reflect"
	"testing"
)

func TestProblem1299(t *testing.T) {
	answer := []int{18,6,6,6,1,-1}
	result := replaceElements([]int{17,18,5,4,6,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}
	