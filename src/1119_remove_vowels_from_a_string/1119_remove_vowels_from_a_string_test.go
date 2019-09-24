package problem1119

import (
	"reflect"
	"testing"
)

func TestProblem1119(t *testing.T) {
	a := "leetcodeisacommunityforcoders"
	answera := "ltcdscmmntyfrcdrs"
	resulta := removeVowels(a)
	if !reflect.DeepEqual(answera, resulta) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    b := "aeiou"
	answerb := ""
	resultb := removeVowels(b)
	if !reflect.DeepEqual(answerb, resultb) {
		t.Errorf("wrong") // to indicate test failed
	}
}
