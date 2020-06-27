package problem1342

func numberOfSteps (num int) int {
	step := 0
	for num > 0{
		if num % 2 ==0 {
			num /= 2
		}else{
			num--;
		}
		step++
	}
	return step
}