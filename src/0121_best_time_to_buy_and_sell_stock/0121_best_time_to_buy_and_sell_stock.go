package problem0121

func maxProfit(prices []int) int {    
    if len(prices) <= 1{
        return 0
    }
                                  
    min,max,currentProfit := prices[0],-1,0
    
    for _,v := range prices{
        if v < min{
            min = v
        }else{
            currentProfit = v - min
            if currentProfit > max{
                max = currentProfit
            }
        }
    }
    return max
}