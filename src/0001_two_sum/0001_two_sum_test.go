package problem0001

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem0001(t *testing.T) {
	a := []int{2, 7, 11, 15}
	answer := []int{0, 1}
	result := twoSum(a, 9)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
	fmt.Println("the end")
}
