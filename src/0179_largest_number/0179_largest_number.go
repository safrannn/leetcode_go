package problem0179

import ("strconv"
		"sort"
	)


func largestNumber(nums []int) string {
    result :=  ""
    sort.Slice(nums, func(i, j int) bool{
        first, second := strconv.Itoa(nums[i]),strconv.Itoa(nums[j])
        if first + second >= second + first{
            return true
        }
        return false
    })
       
   for _,v := range nums{
	   result += strconv.Itoa(v)
   }
   if result[0] == '0'{
	   return result[:1]
   }
   return result
}