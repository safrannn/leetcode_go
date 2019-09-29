package problem1165


import (
	"reflect"
	"testing"
)

func TestProblem1165(t *testing.T) {
	answer := 16
	result := calculateTime("pqrstuvwxyzabcdefghijklmno","f")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}
	