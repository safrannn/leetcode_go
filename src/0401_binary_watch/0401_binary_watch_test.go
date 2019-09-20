package problem0401

import (
	"reflect"
	"testing"
)


func TestProblem0401(t *testing.T) {
	a := 0
	answera := []string{"0:00"}
	resulta := readBinaryWatch(a)
	if !reflect.DeepEqual(answera, resulta) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    b := 1
	answerb := []string{"0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"}
	resultb := readBinaryWatch(b)
	if !reflect.DeepEqual(answerb, resultb) {
		t.Errorf("wrong") // to indicate test failed
    }
    
}