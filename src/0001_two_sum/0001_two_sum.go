package problem0001

func twoSum(nums []int, target int) []int {
	aMap := make(map[int]int)

	for i, v := range nums {
		j, w := aMap[target-v]
		if w {
			return []int{j, i}
		}
		aMap[v] = i
	}
	return []int{}
}
