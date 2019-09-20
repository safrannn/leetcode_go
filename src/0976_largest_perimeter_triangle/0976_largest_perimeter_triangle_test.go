package problem0976

import (
	"reflect"
	"testing"
)


func TestProblem0976(t *testing.T) {
	a := []int{2,1,2}
	answera := 5
	resulta := largestPerimeter(a)
	if !reflect.DeepEqual(answera, resulta) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    b:= []int{1,2,1}
	answerb := 0
	resultb := largestPerimeter(b)
	if !reflect.DeepEqual(answerb, resultb) {
		t.Errorf("wrong") // to indicate test failed
    }
        
    c := []int{3,2,3,4}
	answerc := 10
	resultc := largestPerimeter(c)
	if !reflect.DeepEqual(answerc, resultc) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    d := []int{3,6,2,3}
	answerd := 8
	resultd := largestPerimeter(d)
	if !reflect.DeepEqual(answerd, resultd) {
		t.Errorf("wrong") // to indicate test failed
	}
}
