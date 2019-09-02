package problem0153

func findMin(nums []int) int {
	left, right, mid := 0, len(nums)-1, 0

	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid = (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
