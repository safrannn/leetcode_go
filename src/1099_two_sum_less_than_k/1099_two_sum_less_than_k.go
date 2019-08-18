package problem1099

import "sort"

func twoSumLessThanK(A []int, K int) int {
	if len(A) < 2 {
		return -1
	}
	i, j := 0, len(A)-1
	max := -1
	sort.Ints(A)
	for i < j {
		if A[i]+A[j] < K {
			max = getMax(max, A[i]+A[j])
			i++
		} else {
			j--
		}
	}
	return max
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
