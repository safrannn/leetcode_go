package problem0670

import ("strconv")

func maximumSwap(num int) int {
    numS := []byte(strconv.Itoa(num))
    
    last := make([]int, 10)
    for i := range numS{
        last[numS[i] - '0'] = i
    }
    
    for i := range numS{
        for d := 9; d > int(numS[i] - '0'); d--{
            if last[d] > i{
                numS[i], numS[last[d]] = numS[last[d]], numS[i]
                result,_ := strconv.Atoi(string(numS))
                return result
            }
        }
    }
    return num
}