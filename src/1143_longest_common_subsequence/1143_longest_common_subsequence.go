package problem1143

func longestCommonSubsequence(text1 string, text2 string) int {
    memoMatrix := make([][]int, len(text1)+1)
    for i := range memoMatrix{
        memoMatrix[i] = make([]int, len(text2)+1)
    }
    
    for i := 0; i < len(text1); i++{
        for j := 0; j < len(text2); j++{
            if text1[i] == text2[j]{
                memoMatrix[i+1][j+1] = memoMatrix[i][j] + 1
            }else{
                memoMatrix[i+1][j+1] = max(memoMatrix[i+1][j],memoMatrix[i][j+1])
            }
        }
    }
    return memoMatrix[len(memoMatrix)-1][len(memoMatrix[0])-1]
}
func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}