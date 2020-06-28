package problem1275

import (
	"reflect"
	"testing"
)

func TestProblem0456(t *testing.T) {
	if !reflect.DeepEqual(tictactoe([][]int{[]int{0,0},[]int{2,0},[]int{1,1},[]int{2,1},[]int{2,2}}), "A") {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(tictactoe([][]int{[]int{0,0},[]int{1,1},[]int{0,1},[]int{0,2},[]int{1,0},[]int{2,0}}), "B") {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(tictactoe([][]int{[]int{0,0},[]int{1,1},[]int{2,0},[]int{1,0},[]int{1,2},[]int{2,1},[]int{0,1},[]int{0,2},[]int{2,2}}), "Draw") {
		t.Errorf("wrong") // to indicate test failed
	}
	if !reflect.DeepEqual(tictactoe([][]int{[]int{0,0},[]int{1,1},}), "Pending") {
		t.Errorf("wrong") // to indicate test failed
	}
}