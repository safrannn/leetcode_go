package problem0645

import (
	"reflect"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	a := []int{1, 2, 2, 4}
	answer := []int{2, 3}
	result := findErrorNums(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
