package problem1130

func mctFromLeafValues(arr []int) int {
    result := 0
    stack := []int{}
    stack = append(stack, 2147483648)
    for _,v := range arr{
        for stack[len(stack)-1] <= v{
            current := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result +=  current*min(stack[len(stack)-1], v)
        }
        stack = append(stack,v)
    }
    
    for i := len(stack)-1; i > 1;i--{
        result += stack[i] * stack[i-1]
    } 
    
    return result
}

func min(a,b int)int{
    if a < b{
        return a
    }
    return b
}