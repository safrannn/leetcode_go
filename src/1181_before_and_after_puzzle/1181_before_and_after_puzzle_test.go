package problem1181
import (
	"reflect"
    "testing"
)

func TestProblem1181(t *testing.T) {
	answer := []string{"writing code rocks"}
	result := beforeAndAfterPuzzles([]string{"writing code","code rocks"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = []string{"a"}
	result = beforeAndAfterPuzzles([]string{"a","b","a"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = []string{}
	result = beforeAndAfterPuzzles([]string{"mission complete"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}