package problem0006

func convert(s string, numRows int) string {
    if numRows == 1{
        return s
    }
    letterMatrix := make([][]rune, numRows)
    
    helper,helperSign := 0,-1
    
    for _,v := range s{
        letterMatrix[helper] = append(letterMatrix[helper],v)
        if helper == 0 || helper == numRows -1 {
            helperSign *= -1
        }
        helper += helperSign
    }
    result := ""
    for _,v := range letterMatrix{
        result += string(v)
    }
    return result
}