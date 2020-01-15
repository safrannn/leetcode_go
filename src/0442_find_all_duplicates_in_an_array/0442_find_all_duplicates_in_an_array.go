package problem0442

func findDuplicates(nums []int) []int {
    result := []int{}
    for i := 0; i < len(nums); i++{
        index := 0
        if nums[i] < 0{
            index = -nums[i] - 1
        }else{
            index = nums[i] - 1
        }
        
        if nums[index] < 0{
            result = append(result,index+1)
        }
        
        nums[index] *= -1
    }
    return result
}
