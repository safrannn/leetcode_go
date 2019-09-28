package problem1170

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
    result := make([]int,len(queries))
    wordsCount := make([]int, len(words))
    for i,v := range words{
        wordsCount[i] = getSmallest(v)
    }
    sort.Ints(wordsCount)
    for i,v := range queries{
        fQuery := getSmallest(v)
        result[i] = len(words) - sort.Search(len(words), func(i int) bool { return fQuery < wordsCount[i] })
    }
    return result
}

func getSmallest(s string)int{
    char := 27
    charCount := 0
    for _,v := range s{
        if int(v - 'a') < char{
            charCount = 1
            char = int(v - 'a')
        }else if int(v - 'a') == char{
            charCount++
        }
    }
    return charCount
}