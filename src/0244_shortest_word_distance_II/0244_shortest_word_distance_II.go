package problem0244

type WordDistance struct {
    wordMap map[string][]int
}

func Constructor(words []string) WordDistance {
    m := map[string][]int{}
    wordDistance := WordDistance{m}
    for i, word := range words {
        m[word] = append(m[word], i)
    }
    return wordDistance
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    list1, list2 := this.wordMap[word1], this.wordMap[word2]
    result := 1<<31 - 1
    i,j := 0,0
    for i < len(list1) && j < len(list2){
        index1, index2:= list1[i], list2[j]
        if index1 < index2{
            result = min(result, index2 - index1)
            i++
        }else{
            result = min(result, index1 - index2)
            j++
        }
    }
    return result
}

func min(a,b int)int{
    if a < b{
        return a
    }
    return b
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */