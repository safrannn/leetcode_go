package problem0287

func findDuplicate(nums []int) int {
	//Floyd's Tortoise and Hare
	//each number in nums is between 1 and n that will point to an index that exist
	//then reduce the problem to a cycle detection
	slow := nums[0]
	fast := nums[slow]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	//find entrance to the cycle
	fast = 0

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
