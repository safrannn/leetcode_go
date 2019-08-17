package _1_Two_Sum

import "testing"

func TestTwoSum(t *testing.T) {
	a := []int{2,7,11,15}
	b := []int{0,1}
	
	if twoSum(a,9) != b{
		t.Errorf("wrong") 
	}else{
		fmt.Println("right")
	}
}