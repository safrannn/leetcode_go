package problem0721

import (
	"reflect"
	"testing"
)

func TestProblem0721(t *testing.T) {
	answer := [][]string{[]string{"John", "john00@mail.com", "john_newyork@mail.com","johnsmith@mail.com"},
	[]string{"John", "johnnybravo@mail.com"},
	[]string{"Mary", "mary@mail.com"}}
	result := accountsMerge([][]string{[]string{"John", "johnsmith@mail.com", "john00@mail.com"},
	[]string{"John", "johnnybravo@mail.com"},
	[]string{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
	[]string{"Mary", "mary@mail.com"}})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

}