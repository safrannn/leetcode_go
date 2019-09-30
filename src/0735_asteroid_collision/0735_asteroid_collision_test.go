package problem0735

import (
	"reflect"
	"testing"
)

func TestProblem0455(t *testing.T) {
	answer := []int{5,10}
	result := asteroidCollision([]int{5,10,-5})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = []int{}
	result = asteroidCollision([]int{8,-8})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = []int{-2,-1,1,2}
	result = asteroidCollision([]int{-2,-1,1,2})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}

