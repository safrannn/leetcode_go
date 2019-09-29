package problem1200

import "sort"

func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    smallestDifference := arr[len(arr)-1] + 1
    result := [][]int{}
    for i := 0; i < len(arr) - 1; i++{
        if arr[i+1] - arr[i] == smallestDifference{
            result = append(result, []int{arr[i], arr[i+1]})
        }else if arr[i+1] - arr[i] < smallestDifference{
            result = [][]int{[]int{arr[i],arr[i+1]}}
            smallestDifference = arr[i+1] - arr[i] 
        }
    }
    return result
}