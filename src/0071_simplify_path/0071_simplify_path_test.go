package problem0071

import (
	"reflect"
	"testing"
)


func TestProblem0071(t *testing.T) {
	answer := "/home"
	result := simplifyPath("/home/")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}

	answer = "/"
	result = simplifyPath( "/../")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}

	answer = "/c"
	result = simplifyPath("/a/./b/../../c/")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
