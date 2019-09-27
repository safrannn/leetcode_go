package problem1189

import (
	"reflect"
	"testing"
)

func TestProblem1189(t *testing.T) {
	answer := 1
	result := maxNumberOfBalloons("nlaebolko")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 2
	result = maxNumberOfBalloons("loonbalxballpoon")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    answer = 0
	result = maxNumberOfBalloons("leetcode")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}