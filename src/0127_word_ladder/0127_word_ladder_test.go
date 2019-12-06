package problem0127

import (
	"reflect"
	"testing"
)

func TestProblem0127(t *testing.T) {
	answer := 5
	result := ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log","cog"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
	
	answer = 0
	result = ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
	
	answer = 0
	result = ladderLength("hot", "dog", []string{"hot","dog"})
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
    }
}