package problem1196

import (
	"reflect"
	"testing"
)

func TestProblem1196(t *testing.T) {
	answer := 4
	result := maxNumberOfApples([]int{100,200,150,1000})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 5
	result = maxNumberOfApples([]int{900,950,800,1000,700,800})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}