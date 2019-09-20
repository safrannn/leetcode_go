package problem0191

import (
    "reflect"
    "testing"
)


func TestProblem0142(t *testing.T) {
    var a uint32
    a = 00000000000000000000000000001011
    answera := 3
    resulta := hammingWeight(a)
    if !reflect.DeepEqual(answera, resulta) {
        t.Errorf("wrong") // to indicate test failed
    }

    var b uint32
    b = 00000000000000000000000010000000
    answerb := 1
    resultb := hammingWeight(b)
    if !reflect.DeepEqual(answerb, resultb) {
        t.Errorf("wrong") // to indicate test failed
    }
}