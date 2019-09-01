package problem0162

// Recursive Binary Search
// func findPeakElement(nums []int) int {
//     var search func(left int, right int)int
//     search = func(left,right int)int{
//         if left == right{
//             return left
//         }
//         mid := (left + right) / 2
//         if nums[mid] > nums[mid+1]{
//             return search(left, mid)
//         }else{
//             return search(mid+1, right)
//         }
//     }
//     return search(0,len(nums)-1)
// }

// Iterative
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
