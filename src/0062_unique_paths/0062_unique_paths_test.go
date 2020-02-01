package problem0062
import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	answer :=  3
	result := uniquePaths(3,2)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer =  28
	result = uniquePaths(7,3)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}