package problem0039

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	for i,v := range candidates{
		if v < target{
			for _,w := range combinationSum(candidates[i:], target - v){
				temp := append(w, v)
				result = append(result,temp)
			}
		}else if v == target{
			result = append(result, []int{v})
		}
	}

	return result
}