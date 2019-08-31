package problem0961

func repeatedNTimes(A []int) int {
	existMap := make(map[int]struct{})
	for _, v := range A {
		_, prs := existMap[v]
		if prs {
			return v
		} else {
			existMap[v] = struct{}{}
		}
	}
	return 0
}
