package problem0122

import (
	"reflect"
	"testing"
)

func TestProblem0121(t *testing.T) {
	answer := 7
	result := maxProfit([]int{7,1,5,3,6,4})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 4
	result = maxProfit([]int{1,2,3,4,5})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 0
	result = maxProfit([]int{7,6,4,3,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}