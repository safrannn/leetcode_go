package problem0078

func subsets(nums []int) [][]int {
    result := make([][]int,1) // need to use make instead of initialize as [][]int{}
    for _,v := range nums{
        for _,w := range result{
            result = append(result, append([]int{v},w...))
        }
    }
    return result
}