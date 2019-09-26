package problem0276

func numWays(n int, k int) int {
    if n == 0{
        return 0
    }else if n == 1{
        return k
    }
    //different or the same color with previous
    diff, same := k * (k-1), k
    
    for i := 2; i < n; i++{
        //if the previous two are not of the same color, both color are possible
        diff, same = (diff + same) * (k-1), diff
    }
    
    return diff + same
}