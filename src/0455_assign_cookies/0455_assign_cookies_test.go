package problem0455
 
import (
	"reflect"
	"testing"
)

func TestProblem0455(t *testing.T) {
	answer := 1
	result := findContentChildren([]int{1,2,3}, []int{1,1})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    answer = 2
	result = findContentChildren([]int{1,2}, []int{1,2,3})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}