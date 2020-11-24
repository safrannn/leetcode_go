package problem0213

func rob(nums []int) int {
    n := len(nums);
    if n == 0{
        return 0
    }else if n == 1{
        return nums[0]
    }
    return max(dp(0,n-1, nums), dp(1,n,nums));
} 

func dp(start int, end int, nums []int) int{
    prev, current := 0,0;
    
    for i := start; i < end; i++{
        temp := current;
        current = max(current, nums[i] + prev);
        prev = temp;
    }
    
    return current;
}


func max(a,b int)int{
    if a > b{
        return a
    }else{
        return b
    }
}
