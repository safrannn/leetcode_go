package problem0447

import (
	"reflect"
	"testing"
)


func TestProblem0447(t *testing.T) {
	a := [][]int{[]int{0,0}, []int{1,0}, []int{2,0}}
	answer := 2
	result := numberOfBoomerangs(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}


