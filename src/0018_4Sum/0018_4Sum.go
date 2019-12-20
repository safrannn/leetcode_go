package problem0018

import "sort"

func fourSum(nums []int, target int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums)-3;i++{
        if i != 0 && nums[i] == nums[i-1]{
            continue
        }
        if nums[i] + nums[i+1]+ nums[i+2]+ nums[i+3] > target{
            break
        }
        if nums[i] + nums[len(nums)-1]+ nums[len(nums)-2]+ nums[len(nums)-3] < target{
            continue
        }
        for j := i + 1; j < len(nums)-2; j++{
            if j != i + 1 && nums[j] == nums[j-1]{
                continue
            }
            if nums[i] + nums[j]+ nums[j+1]+ nums[j+2] > target{
                break
            }
            if nums[i] + nums[j]+ nums[len(nums)-1]+ nums[len(nums)-2] < target{
                continue
            }
            left, right := j + 1, len(nums)-1
            for left < right{
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                if sum == target{
                    result = append(result, []int{nums[i], nums[j],nums[left],nums[right]})
                    left++

                    for left < right &&  nums[left] == nums[left-1]{
                        left++
                    }
                }else if sum < target{
                    left++
                }else if sum > target{
                    right--
                }
            }
            
        }
    }
    return result
}