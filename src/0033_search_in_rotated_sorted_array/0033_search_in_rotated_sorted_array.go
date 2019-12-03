package problem0033

func search(nums []int, target int) int {
	left,right := 0, len(nums) - 1

	for left <= right{
		medium := (left + right) / 2

		if nums[medium] == target{
			return medium
		}else if nums[medium] >= nums[left]{
			if nums[left] <= target && target < nums[medium]{
				right = medium - 1
			}else{
				left = medium + 1
			}
		}else{
			if nums[medium] < target && target <= nums[right]{
				left = medium + 1
			}else{
				right = medium - 1
			}
		}
	}
	
	return -1
}