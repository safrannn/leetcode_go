package problem0442

import (
    "reflect"
    "testing"
)

func TestProblem0442(t *testing.T) {
    answer :=  []int{2,3}
    result := findDuplicates([]int{4,3,2,7,8,2,3,1})
    if !reflect.DeepEqual(answer, result) {
        t.Errorf("wrong") // to indicate test failed
    }
}
