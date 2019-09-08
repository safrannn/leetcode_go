package problem0504

import (
	"reflect"
    "testing"
    "fmt"
)


func TestProblem0142(t *testing.T) {
	a := 100
	answer := "202"
	result := convertToBase7(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}
