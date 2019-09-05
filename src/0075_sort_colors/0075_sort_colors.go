package problem0075

func sortColors(nums []int) []int {
	leftBound, rightBound := 0, len(nums)-1

	for i := 0; i <= rightBound; i++ {
		if nums[i] == 0 {
			nums[i], nums[leftBound] = nums[leftBound], nums[i]
			leftBound++
		} else if nums[i] == 2 {
			nums[i], nums[rightBound] = nums[rightBound], nums[i]
			rightBound--
			i--
		}
	}
}
