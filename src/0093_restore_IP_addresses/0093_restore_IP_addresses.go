package problem0093

import "strconv"

func restoreIpAddresses(s string) []string {
    result := []string{}
    
    if len(s) > 12 {
        return result
    }
    
    var backtrack func(string, string, int)
    backtrack = func(previous string, remain string, dotUsed int){
        if dotUsed == 4 && len(remain) == 0{
            result = append(result, previous)
            return
        }
        
        if len(previous) != 0{
            previous += "."
        }
        
        //get current block of number
        for i := 1; i < 4; i++{
            // break if no more numbers
            if len(remain) < i{
                break
            }
            // break if leading zero, 0 only is allowed
            if i > 1 && remain[0] == '0'{
                break
            }
            
            // for every possible number, add it to previous and dfs backtracking
            value, _ := strconv.Atoi(remain[:i])
            if value < 256{
                backtrack(previous + remain[:i], remain[i:], dotUsed+1)
            }            
        }
        
    }
    
    backtrack("", s, 0)
    
    return result
}