package problem0935

func knightDialer(N int) int {
    mod := 1000000007
    neighbor := make(map[int][]int,10)
    neighbor[0] = []int{4,6}
    neighbor[1] = []int{6,8}
    neighbor[2] = []int{7,9}
    neighbor[3] = []int{4,8}
    neighbor[4] = []int{3,9,0}
    neighbor[5] = []int{}
    neighbor[6] = []int{1,7,0}
    neighbor[7] = []int{2,6}
    neighbor[8] = []int{1,3}
    neighbor[9] = []int{2,4}
    
    dp := []int{1,1,1,1,1,1,1,1,1,1}
    
    for N > 1{
        dp2 := make([]int,10)
        for i,v := range dp{
            for _,neighb := range neighbor[i]{
                dp2[neighb] += v
                dp2[neighb] %= mod
            }
        }
        dp = dp2
        N--
    }
    
    result := 0
    for _,v := range dp{
        result += v 
    }
    
    return result % mod
    
}