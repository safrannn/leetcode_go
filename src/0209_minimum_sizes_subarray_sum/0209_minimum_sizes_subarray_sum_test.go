package problem0209


import (
    "reflect"
    "testing"
)

func TestProblem0209(t *testing.T) {
    answer := 2
    result := minSubArrayLen(7,[]int{2,3,1,2,4,3})
    if !reflect.DeepEqual(answer, result) {
        t.Errorf("wrong") // to indicate test failed
    }
}