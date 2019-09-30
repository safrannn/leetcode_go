package problem1196

import "sort"

func maxNumberOfApples(arr []int) int {
    sort.Ints(arr)
    count,sum := 0, 0
    for _,v := range arr{
        sum += v
        if sum > 5000{
            break
        }
        count++
    }
    return count
}