package problem1232
import (
	"reflect"
	"testing"
)

func TestProblem1207(t *testing.T) {
	answer := true
	result := checkStraightLine([][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{4,5},[]int{5,6},[]int{6,7}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }


}
