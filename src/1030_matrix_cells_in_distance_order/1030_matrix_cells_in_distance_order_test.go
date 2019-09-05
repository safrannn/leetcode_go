package problem1030

import (
	"reflect"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	aanswer := [][]int{[]int{0, 1}, []int{0, 0}, []int{1, 0}, []int{1, 1}}
	result := allCellsDistOrder(2, 2, 0, 1)
	if !reflect.DeepEqual(aanswer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
