package problem0961

import (
	"reflect"
	"testing"
)

func Testroblem0961(t *testing.T) {
	a1 := []int{1, 2, 3, 3}
	a1answer := 3
	result1 := repeatedNTimes(a1)
	if !reflect.DeepEqual(a1answer, result1) {
		t.Errorf("wrong") // to indicate test failed
	}

	a2 := []int{2, 1, 2, 5, 3, 2}
	a2answer := 2
	result2 := repeatedNTimes(a2)
	if !reflect.DeepEqual(a2answer, result2) {
		t.Errorf("wrong") // to indicate test failed
	}

	a3 := []int{5, 1, 5, 2, 5, 3, 5, 4}
	a3answer := 5
	result3 := repeatedNTimes(a3)
	if !reflect.DeepEqual(a3answer, result3) {
		t.Errorf("wrong") // to indicate test failed
	}
}
