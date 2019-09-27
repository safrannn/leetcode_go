package problem1160

func countCharacters(words []string, chars string)


import (
	"reflect"
	"testing"
)

func TestProblem1160(t *testing.T) {
	answer := 6
	result := countCharacters([]string{"cat","bt","hat","tree"},"atach")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }

    answer := 10
	result := countCharacters([]string{"hello","world","leetcode"},"welldonehoneyr")
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}