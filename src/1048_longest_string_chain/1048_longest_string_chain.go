package problem1048

import "sort"

func longestStrChain(words []string) int {
    sort.SliceStable(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    
    wordsMap := make(map[int][]string)
    for _,v := range words{
        wordsMap[len(v)] = append(wordsMap[len(v)],v)
    }
    
    wordsPreCount := make(map[string]int)
    result := 0
    
    for i := 1; i <= 16; i++{
        if len(wordsMap[i]) == 0{
            continue
        }
        
        for _,word := range wordsMap[i]{
            _,prs := wordsPreCount[word]
            if !prs{
                wordsPreCount[word] = 1
            }
            result = max(result, wordsPreCount[word])
            
            for _, newWord := range wordsMap[i+1]{
                if isPre(word,newWord){
                    wordsPreCount[newWord] = max(wordsPreCount[newWord],wordsPreCount[word] + 1)
                    result = max(result,wordsPreCount[newWord] )
                }
            }
        }
        
    }
    return result
    
}

func isPre(cand, word string)bool{
    iCand,iWord := 0,0
    diff := false
    for iCand < len(cand) && iWord < len(word){
        if cand[iCand] == word[iWord]{
            iCand++
            iWord++
        }else{
            if diff {
                return false
            }
            diff = true
            iWord++
        }
    }
    return true
}

func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}