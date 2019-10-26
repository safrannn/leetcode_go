package problem0179


import (
	"reflect"
	"testing"
)

func TestProblem0179(t *testing.T) {
	answer := "210"
	result := largestNumber([]int{10,2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = "9534330"
	result = largestNumber([]int{3,30,34,5,9})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}