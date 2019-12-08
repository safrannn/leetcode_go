package problem1221

func balancedStringSplit(s string) int {
    current := 0
    result := 0
    
    for i := range s{
        if s[i] == 'L'{
            current++
        }else{
            current--
        }
        if current == 0{
            result++
        }
    }
    
    return result
}