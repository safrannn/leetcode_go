package problem1108

import (
	"reflect"
	"testing"
)
func TestProblem1108(t *testing.T) {
	a := "1.1.1.1"
	answera := "1[.]1[.]1[.]1"
	resulta := defangIPaddr(a)
	if !reflect.DeepEqual(answera, resulta) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    b := "255.100.50.0"
	answerb := "255[.]100[.]50[.]0"
	resultb := defangIPaddr(b)
	if !reflect.DeepEqual(answerb, resultb) {
		t.Errorf("wrong") // to indicate test failed
    }
}