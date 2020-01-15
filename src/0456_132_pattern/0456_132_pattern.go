package problem0456

func find132pattern(nums []int) bool {
    if len(nums) < 3{
        return false
    }
    
    // get minimum number appeared before and include nums[i]
    minimums := make([]int,len(nums))
    minimums[0] = nums[0]
    for i := 1; i < len(nums); i++{
        minimums[i] = min(nums[i],minimums[i-1])
    }
    
    // use stack to find number that satisfy the condition
    stack := []int{}
    for i := len(nums)-1; i>=0;i--{
        if nums[i] > minimums[i]{
            for len(stack) > 0 && stack[len(stack)-1] <= minimums[i]{
                stack = stack[:len(stack)-1]
            }
            if len(stack) > 0 && stack[len(stack)-1] < nums[i]{
                return true
            }
            stack = append(stack, nums[i])
        }
    }
    
    return false
}


func min(a,b int)int{
    if a < b {
        return a
    }
    return b
}