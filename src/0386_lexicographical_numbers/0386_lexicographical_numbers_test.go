package problem0386

import (
	"reflect"
	"testing"
)


func TestProblem0386(t *testing.T) {
	answera := []string{1,10,11,12,13,2,3,4,5,6,7,8,9}
	resulta := lexicalOrder([]int{1,10,11,12,13,2,3,4,5,6,7,8,9})
	if !reflect.DeepEqual(answera, resulta) {
		t.Errorf("wrong") // to indicate test failed
    }
}