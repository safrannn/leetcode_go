package problem0075

import (
	"reflect"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	a1 := []int{2, 0, 2, 1, 1, 0}
	a1answer := []int{0, 0, 1, 1, 2, 2}
	result1 := sortColors(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}
}
