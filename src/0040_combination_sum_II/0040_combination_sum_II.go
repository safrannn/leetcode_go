package problem0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)  
    result := [][]int{}
    dfs(candidates, target, 0, []int{}, &result)
    return result
}

func dfs(candidates []int, target int, sum int, currentList []int, result *[][]int){
    for i,v := range candidates{
        if i > 0 && candidates[i] == candidates[i-1]{
            continue
        }
          
        if sum + v < target{
            dfs(candidates[i+1:], target, sum + v, append(currentList,v),result)
        }else if sum + v == target{
            *result = append(*result, append(currentList,v))
        } 
    }
}

