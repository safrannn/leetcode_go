package problem0274

func hIndex(citations []int) int {
    n := len(citations)
    paperCount := make([]int,n+1)
	
	// counting papers for each citation number
    for _,v := range citations{
        paperCount[min(n,v)]++
    }
	
	// finding the h-index
    citationCountSum := paperCount[n]
    h := n
    for h > citationCountSum{
        h--
        citationCountSum += paperCount[h]
    }
    return h
}

func min(a,b int)int{
    if a < b{
        return a
    }
    return b
}