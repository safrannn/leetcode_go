package problem0787

import (
	"reflect"
	"testing"
)

func TestProblem0787(t *testing.T) {
	a1 := [][]int{[]int{0, 1, 100}, []int{1, 2, 100}, []int{0, 2, 500}}
	a1answer := 200
	result1 := findCheapestPrice(3, a1, 0, 2, 1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}
}
