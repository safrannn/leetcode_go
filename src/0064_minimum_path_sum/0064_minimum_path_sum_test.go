package problem0064

import (
    "reflect"
    "testing"
)

func TestProblem0064(t *testing.T) {
    answer := 7
    result := minPathSum([][]int{[]int{1,3,1},[]int{1,5,1},[]int{4,2,1}})
    if !reflect.DeepEqual(answer, result) {
        t.Errorf("wrong") // to indicate test failed
    }
}