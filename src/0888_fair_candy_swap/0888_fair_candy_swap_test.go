package problem0888

import (
	"reflect"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	a1 := []int{1, 1}
	a2 := []int{2, 2}
	aanswer := []int{1, 2}
	result := fairCandySwap(a1, a2)
	if !reflect.DeepEqual(aanswer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
