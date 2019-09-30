package problem1207

func uniqueOccurrences(arr []int) bool {
    occurance := make(map[int]int)
    for _,v := range arr{
        occurance[v]++
    }
    helper := make([]int,len(arr))
    for i := range occurance{
        if helper[occurance[i]] > 0{
            return false
        }
        helper[occurance[i]]++
    }
    return true
}