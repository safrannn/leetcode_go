package problem1249

func minRemoveToMakeValid(s string) string {
    stack := []int{}
    stackCount := 0
    removeIndex := make([]bool,len(s))
    
    for i,v := range s{
        if v == '('{
            stack = append(stack,i)
            stackCount++
        }else if v == ')'{
            if stackCount == 0{
                removeIndex[i] = true
            }else{
                stack = stack[:len(stack)-1]
                stackCount--
            }
        }
    }

    for _,v := range stack{
        removeIndex[v] = true
    }
    result := []byte{}
    for i,v := range removeIndex{
        if !v {
            result = append(result,s[i])
        }
    }
    return string(result)
}