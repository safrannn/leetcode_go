package problem0739

func dailyTemperatures(T []int) []int {
    stack := []int{}
    result := make([]int,len(T))

    for i:= range T{
        for len(stack) > 0 && T[i] > T[stack[len(stack)-1]]{
            result[stack[len(stack) - 1]] = i - stack[len(stack)-1] 
            stack = stack[:len(stack)-1]
        }
        stack = append(stack,i)
    } 
    
    return result
}