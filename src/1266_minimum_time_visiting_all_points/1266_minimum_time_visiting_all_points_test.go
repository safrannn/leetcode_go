package problem1266

import (
	"reflect"
	"testing"
)
    
func TestProblem1266(t *testing.T) {
    answer := 7
    result := minTimeToVisitAllPoints([][]int{[]int{1,1},[]int{3,4},[]int{-1,0}})
    if !reflect.DeepEqual(answer, result) {
        t.Errorf("wrong") // to indicate test failed
    }

    answer = 5
    result = minTimeToVisitAllPoints([][]int{[]int{3,2},[]int{-2,2}})
    if !reflect.DeepEqual(answer, result) {
        t.Errorf("wrong") // to indicate test failed
    }

}