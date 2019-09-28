package problem1170

import (
	"reflect"
	"testing"
)

func TestProblem0455(t *testing.T) {
	answer := []int{6,1,1,2,3,3,3,1,3,2}
	result := numSmallerByFrequency([]string{"bba","abaaaaaa","aaaaaa","bbabbabaab","aba","aa","baab","bbbbbb","aab","bbabbaabb"},[]string{"aaabbb","aab","babbab","babbbb","b","bbbbbbbbab","a","bbbbbbbbbb","baaabbaab","aa"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}