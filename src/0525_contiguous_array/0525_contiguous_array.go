package problem0525

func findMaxLength(nums []int) int {
	helperMap := make(map[int]int)
	helperMap[0] = -1
	count := 0
	max := 0
	for i := range nums{
		if nums[i] == 0{
			count--
		}else{
			count++
		}
		_,prs := helperMap[count]
		if prs{
			max = getMax(max, i - helperMap[count])
		}else{
			helperMap[count] = i
		}
	}
	return max
}

func getMax(a,b int)int{
	if a > b{
		return a 
	}
	return b
}