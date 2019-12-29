package problem0003

func lengthOfLongestSubstring(s string) int {
    index := make([]int, 128)
    result := 0
    for i,j := 0,0; j < len(s); j++{
        i = max(index[s[j]] ,i)
        result = max(result, j - i + 1)
        index[s[j]] = j + 1
    }
    return result
}
func max(a,b int)int{
    if a < b{
        return b
    }
    return a
}