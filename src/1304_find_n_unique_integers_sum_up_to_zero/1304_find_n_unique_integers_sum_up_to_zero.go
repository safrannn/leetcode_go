package problem1304

func sumZero(n int) []int {
	result := make([]int, 0, n)
    if n & 1 == 1{
		result = append(result,0)
	}
	for i := 1; i <= n/2; i++{
		result = append(result,i)
		result = append(result,-i)
	}
	return result
}