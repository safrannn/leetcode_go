package problem0492

import (
	"reflect"
	"testing"
)

func TestProblem0492(t *testing.T) {
	answer := []int{2,2}
	result := constructRectangle(4)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
	
	answer = []int{5,1}
	result = constructRectangle(5)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}