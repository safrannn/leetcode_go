package problem1304

import (
	"reflect"
	"testing"
)

func TestProblem1304(t *testing.T) {
	answer := []int{0,1,-1,2,-2}
	result := sumZero(5)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}
	