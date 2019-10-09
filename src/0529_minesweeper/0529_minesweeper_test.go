package problem0529

import (
	"reflect"
	"testing"
)

func TestProblem0529(t *testing.T) {
	answer := [][]byte{[]byte{'B','1','E','1','B'},[]byte{'B','1','M','1','B'},[]byte{'B','1','1','1','B'},[]byte{'B','B','B','B','B'}}
	result := updateBoard([][]byte{[]byte{'E','E','E','E','E'},[]byte{'E','E','M','E','E'},[]byte{'E','E','E','E','E'},[]byte{'E','E','E','E','E'}},[]int{3,0})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}