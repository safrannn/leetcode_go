package problem1323

import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	answer :=  9969
	result := maximum69Number(9669)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}