package problem0008

import ("math"
"strings")

func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    if len(str) == 0 {
        return 0
    }
    sign := 1
    i:=0
    answer := 0
    if str[0] == '-'{
        sign *= -1
        i++
    }else if str[0] == '+'{
        i++
    }
    
    for i < len(str){
        if str[i] < '0' || str[i] > '9'{
            return answer*sign
        }else{
            answer =  int(str[i] - '0')+answer*10
            if answer > math.MaxInt32{
                if sign == 1{
                    return math.MaxInt32
                }else{
                    return math.MinInt32
                }
            }
            i++
        }
    }
    
    return answer*sign
}