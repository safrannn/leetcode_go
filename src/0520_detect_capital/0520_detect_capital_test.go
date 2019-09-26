package problem0520

import (
	"reflect"
	"testing"
)

func TestProblem0455(t *testing.T) {
	answer := true
	result := detectCapitalUse("USA")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    answer = false
	result = detectCapitalUse("FlaG")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
    
    answer = false
	result = detectCapitalUse("FFFFFFf")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}