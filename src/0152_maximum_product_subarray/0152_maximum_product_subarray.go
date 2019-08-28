package problem0152

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	result, min, max := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}

		max *= nums[i]
		if max < nums[i] {
			max = nums[i]
		}

		min *= nums[i]
		if min > nums[i] {
			min = nums[i]
		}

		if max > result {
			result = max
		}
	}
	return result
}
