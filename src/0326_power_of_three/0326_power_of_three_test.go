package problem0326

import (
	"reflect"
	"testing"
)

func TestProblem0326(t *testing.T) {
	a1 := 27
	a1answer := true
	a1result := isPowerOfThree(a1)
	if !reflect.DeepEqual(a1answer, a1result) {
		t.Errorf("wrong") // to indicate test failed
	}

	a2 := 0
	a2answer := false
	a2result := isPowerOfThree(a2)
	if !reflect.DeepEqual(a2answer, a2result) {
		t.Errorf("wrong") // to indicate test failed
	}

	a3 := 45
	a3answer := false
	a3result := isPowerOfThree(a3)
	if !reflect.DeepEqual(a3answer, a3result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
