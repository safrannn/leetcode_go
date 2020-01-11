package problem0209

func minSubArrayLen(s int, nums []int) int {
    sum := 0
    result := len(nums) + 1
    leftIndex := 0
    
    for i := 0; i < len(nums); i++{
        sum += nums[i]
        for sum >= s{
            result = min(result, i+1-leftIndex)
            sum -= nums[leftIndex]
            leftIndex++
        }
    }
    if result == len(nums) + 1{
        return 0
    }
    return result
}

func min(a,b int)int{
    if a < b{
        return a
    }
    return b
}