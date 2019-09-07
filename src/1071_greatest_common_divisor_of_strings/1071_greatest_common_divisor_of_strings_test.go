package problem1071

import (
	"reflect"
	"testing"
)

func TestProblem1071(t *testing.T) {
    astr1 := "ABCABC"
    astr2 := "ABC"
	answera := "ABC"
	resulta := gcdOfStrings(astr1,astr2)
	if !reflect.DeepEqual(answera, resulta) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    bstr1 := "ABABAB"
    bstr2 := "ABAB"
	answerb := "AB"
	resultb := gcdOfStrings(bstr1,bstr2)
	if !reflect.DeepEqual(answerb, resultb) {
		t.Errorf("wrong") // to indicate test failed
	}
}

