package problem0016

import "sort"

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    result := nums[0] + nums[1] + nums[len(nums)-1]
    for i := 0; i < len(nums)-2; i++{
        start, end := i+1, len(nums)-1
        for start < end{
            sum := nums[i] + nums[start] + nums[end]
            if sum > target{
                end--
            }else if sum < target{
                start++
            }else{
                return target
            }
            
            if abs(sum - target) < abs(result - target){
                result = sum
            }
        }
    }
    return result
}
func abs(a int)int{
    if a < 0{
        return -a
    }
    return a 
}