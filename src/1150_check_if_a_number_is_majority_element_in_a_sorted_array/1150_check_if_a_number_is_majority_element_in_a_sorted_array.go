package problem1150

func isMajorityElement(nums []int, target int) bool {
    return nums[(len(nums)-1)/2] == target
}