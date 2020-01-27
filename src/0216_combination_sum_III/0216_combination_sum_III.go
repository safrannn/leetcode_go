package problem0216

func combinationSum3(k int, n int) [][]int {
    result := [][]int{}
    combination := []int{1,2,3,4,5,6,7,8,9}
    dfs(&combination, 0, k, n,[]int{}, &result)
    return result
}

func dfs(combination *[]int, start int, k int, n int, list []int, result *[][]int){
    if k == 0 && n == 0{
        newList := make([]int,len(list))
        copy(newList,list)
        *result = append(*result, newList)    
    }else{
        for i := start; i < 9 && k > 0 && n > 0 ; i++{
            v := (*combination)[i]
            dfs(combination, i + 1, k - 1, n - v, append(list,v), result)
        }
    }
}