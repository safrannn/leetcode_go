package problem1029

import ("reflect"
        "testing")


func TestProblem1029(t *testing.T) {
	a := [][]int{[]int{10,20},[]int{30,200},[]int{400,50},[]int{30,20}}
	answer := 110
	result := twoCitySchedCost(a)
	if !reflect.DeepEqual(answer, result) {
		t.Errorf("wrong") // to indicate test failed
	}
}
