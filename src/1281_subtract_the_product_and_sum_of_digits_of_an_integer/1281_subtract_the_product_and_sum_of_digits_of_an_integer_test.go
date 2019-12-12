package problem1281


import (
	"reflect"
	"testing"
)

func TestProblem1207(t *testing.T) {
	answer := 15
	result := subtractProductAndSum(234)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = 21
	result = subtractProductAndSum(4421)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer = -3
	result = subtractProductAndSum(1111)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}
