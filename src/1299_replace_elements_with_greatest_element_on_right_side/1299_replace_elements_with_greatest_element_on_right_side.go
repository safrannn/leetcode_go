package problem1299

func replaceElements(arr []int) []int {
	length := len(arr)
	result := make([]int,length)
	result[length - 1] = -1

    for i := length - 2; i >= 0; i--{
		result[i] = max(arr[i + 1], result[i + 1])
	}
	
	return result
}

func max( a, b int)int{
    if a > b{
        return a
    }
    return b
}