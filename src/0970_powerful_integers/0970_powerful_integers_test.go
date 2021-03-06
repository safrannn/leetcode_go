package problem0970

import (
	"reflect"
	"testing"
)

func TestProblem0970(t *testing.T) {
	answer := []int{2,3,5,9,17,33,65,129,257,513,1025,2049,4097,8193,16385,32769,65537,131073,262145,524289}
	result := powerfulIntegers(1,2,1000000)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
