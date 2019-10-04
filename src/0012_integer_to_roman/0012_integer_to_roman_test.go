package problem0012

import (
	"reflect"
	"testing"
)

func TestProblem0012(t *testing.T) {
	answer := "MMCCCXLIX"
	result := intToRoman(2349)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = "CCCLXXX"
	result = intToRoman(380)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}