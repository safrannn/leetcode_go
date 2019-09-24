package problem1185

import (
	"reflect"
	"testing"
)

func TestProblem1185(t *testing.T) {
	answera := "Saturday"
	resulta := dayOfTheWeek(31,8,2019)
	if !reflect.DeepEqual(answera, resulta) {
		t.Errorf("wrong") // to indicate test failed
    }

	answerb := "Sunday"
	resultb := dayOfTheWeek(18,7,1999)
	if !reflect.DeepEqual(answerb, resultb) {
		t.Errorf("wrong") // to indicate test failed
    }

	answerc := "Sunday"
	resultc := dayOfTheWeek(15,8,1993)
	if !reflect.DeepEqual(answerc, resultc) {
		t.Errorf("wrong") // to indicate test failed
	}
}
