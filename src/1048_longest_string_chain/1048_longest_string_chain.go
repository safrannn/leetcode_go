package problem1048

import "sort"

func longestStrChain(words []string) int {
    // sort words from shortest to longest
    sort.SliceStable(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    // for each word, loops all its possible predecessor
    // update the longest chain length
    result := 0
    helperMap := make(map[string]int)
    for _,v := range words{
        best := 0
        for j := range v{
            prevWord := v[0:j] + v[j+1:]
            best = max(best, helperMap[prevWord] + 1)
        }
        helperMap[v] = best
        result = max(result, best)
    }
    return result
}

func max(a,b int)int{
    if a > b {
        return a 
    }
    return b
}