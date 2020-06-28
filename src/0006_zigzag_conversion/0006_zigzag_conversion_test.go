package problem0006

import (
	"reflect"
	"testing"
)

func TestProblem0006(t *testing.T) {
	if !reflect.DeepEqual(convert("PAYPALISHIRING",3), "PAHNAPLSIIGYIR") {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(convert("PAYPALISHIRING",4), "PINALSIGYAHRPI") {
		t.Errorf("wrong") // to indicate test failed
	}
}