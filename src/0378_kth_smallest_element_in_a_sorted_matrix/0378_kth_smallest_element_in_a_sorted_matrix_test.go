package problem0378

import (
	"reflect"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	a1 := [][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}
	k1 := 8
	a1answer := 13
	result1 := kthSmallest(a1, k1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}
}
