package problem1184

import (
	"reflect"
	"testing"
)

func TestProblem1184(t *testing.T) {
	answer := 1
	result := distanceBetweenBusStops([]int{1,2,3,4},0,1)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 3
	result = distanceBetweenBusStops([]int{1,2,3,4},0,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 4
	result = distanceBetweenBusStops([]int{1,2,3,4},0,3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}