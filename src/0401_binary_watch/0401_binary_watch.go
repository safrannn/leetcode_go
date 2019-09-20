package problem0401

import ("fmt")

func readBinaryWatch(num int) []string {
    result := []string{}
    
    if num < 0{
        return result
    }
    
    for h:=0; h < 12; h++{
        for m :=0; m < 60; m++{
            if countBit(h) + countBit(m) == num{
                s := fmt.Sprintf("%d:%02d", h, m)
                result = append(result, s)
            }
        }
    }
    
    return result
}

func countBit(num int)int{
    result := 0 
    for num > 0{
        if num & 1 == 1{
            result++
        }
        num >>= 1
    }
    return result
}