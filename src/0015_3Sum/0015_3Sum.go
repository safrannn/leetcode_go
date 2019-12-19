package problem0015

import "sort"

func threeSum(nums []int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++{
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        left, right := i+1, len(nums)-1
        target:= -nums[i]
        for left < right{
            sum := nums[left] + nums[right]
            if sum < target{
                left++
            }else if sum > target{
                right--
            }else{
                result = append(result, []int{nums[i], nums[left],nums[right]})
                left++
                right--
                for left < right && nums[left] == nums[left-1]{
                    left++
                }
                for left < right && nums[right] == nums[right+1]{
                    right--
                }
            }
        }
    }
    return result
}