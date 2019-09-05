package problem0605

import (
	"reflect"
	"testing"
)

func TestProblem0605(t *testing.T) {
	a := []int{1, 0, 0, 0, 1}
	answer := false
	result := canPlaceFlowers(a, 2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
