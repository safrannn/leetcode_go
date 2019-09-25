package problem0884

import "strings"

func uncommonFromSentences(A string, B string) []string {
    aSplit := strings.Split(A, " ")
    bSplit := strings.Split(B, " ")
    helperMap := make(map[string]int, len(A)+len(B))
    result := []string{}
    
    for _,v := range aSplit{
        helperMap[v]++
    }
    for _,v := range bSplit{
        helperMap[v]++
    }
    for i := range helperMap{
        if helperMap[i] == 1{
            result = append(result,i)
        }
    }
    return result
}