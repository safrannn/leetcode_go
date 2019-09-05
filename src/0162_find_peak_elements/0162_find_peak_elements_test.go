package problem0162

import (
	"reflect"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	a1 := []int{1, 2, 1, 3, 5, 6, 4}
	a1answer := 5
	result1 := findPeakElement(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}
}
