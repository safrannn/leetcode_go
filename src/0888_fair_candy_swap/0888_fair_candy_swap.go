package problem0888

func fairCandySwap(A []int, B []int) []int {
	candyA := make(map[int]struct{})
	sumA, sumB := 0, 0
	for _, v := range A {
		candyA[v] = struct{}{}
		sumA += v
	}

	for _, v := range B {
		sumB += v
	}

	result := []int{0, 0}
	for _, v := range B {
		_, prs := candyA[(sumA-sumB)/2+v]
		if prs {
			result = []int{(sumA-sumB)/2 + v, v}
		}
	}
	return result
}
