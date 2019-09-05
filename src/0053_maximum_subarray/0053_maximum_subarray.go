package problem0053

func maxSubArray(nums []int) int {
    //greedy
    currentSum := nums[0]
    maxSum := nums[0]
    
    for i:=1; i<len(nums);i++{
        currentSum = getMax(nums[i], currentSum + nums[i])
        maxSum = getMax(maxSum, currentSum)
    } 
    return maxSum
}

func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}