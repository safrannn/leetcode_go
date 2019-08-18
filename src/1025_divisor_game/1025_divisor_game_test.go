package problem1025

import (
	"reflect"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	a := 3
	answer := false
	result := divisorGame(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
