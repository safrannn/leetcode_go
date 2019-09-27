package problem1160

func countCharacters(words []string, chars string) int {
    charsCount := make([]int,26)
    for _,v:= range chars{
        charsCount[v-'a']++
    }
    
    result := 0
    for _,v := range words{
        charsCountCur := make([]int,26)
        
        isGood := true
        for _,w := range v{
            charsCountCur[w-'a']++
            if charsCountCur[w-'a'] > charsCount[w-'a']{
                isGood = false
                break
            }
        }
        if isGood{
            result += len(v)
        }
    }
    return result
}