package problem0645

func findErrorNums(nums []int) []int {
	helper := make([]bool, len(nums))
	result := []int{}
	for _, v := range nums {
		if helper[v-1] {
			result = append(result, v)
		} else {
			helper[v-1] = true
		}
	}

	for i, v := range helper {
		if !v {
			result = append(result, i+1)
			break
		}
	}

	return result
}
