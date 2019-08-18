package problem1086

import (
	"reflect"
	"testing"
)

func TestHighFive(t *testing.T) {
	a := [][]int{[]int{1, 46}, []int{1, 93}, []int{1, 94}, []int{1, 36}, []int{1, 71}, []int{2, 4}, []int{2, 68}, []int{2, 63}, []int{2, 80}, []int{2, 27}, []int{3, 79}, []int{3, 35}, []int{3, 95}, []int{3, 56}, []int{3, 35}, []int{4, 31}, []int{4, 32}, []int{4, 82}, []int{4, 42}, []int{4, 97}, []int{6, 0}, []int{6, 0}, []int{6, 1}, []int{6, 2}, []int{6, 3}, []int{6, 2}, []int{6, 1}}
	answer := [][]int{[]int{1, 68}, []int{2, 48}, []int{3, 60}, []int{4, 56}, []int{6, 1}}
	result := highFive(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
