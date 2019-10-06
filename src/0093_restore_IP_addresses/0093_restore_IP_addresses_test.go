package problem0093

import (
	"reflect"
	"testing"
)

func TestProblem0093(t *testing.T) {
	answer := []string{"255.255.11.135", "255.255.111.35"}
	result := restoreIpAddresses("25525511135")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}