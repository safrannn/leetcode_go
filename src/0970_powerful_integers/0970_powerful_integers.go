package problem0970

import "sort"

func powerfulIntegers(x int, y int, bound int) []int {
    result := make([]int,0,128)
    
    if x == 1{
        x = bound + 1
    }
    
    if y == 1{
        y = bound + 1
    }
    
    for i:= 1; i < bound; i *= x{
        for j := 1; i + j <= bound; j *= y{
            result = append(result, i+j)
        }
    }
    
    
    last, j := -1,-1
    sort.Ints(result)
    for i := range result{
        if last != result[i]{
            j++
            result[j], last = result[i], result[i]
        }
    }
    
    
    return result[:j+1]
}
