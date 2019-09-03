package problem0378

func kthSmallest(matrix [][]int, k int) int {
	//https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/discuss/301357/Java-0ms-(added-Python-and-C%2B%2B)%3A-Easy-to-understand-solutions-using-Heap-and-Binary-Search
	m, n := len(matrix), len(matrix[0])
	low, high := matrix[0][0], matrix[m-1][n-1]+1

	for low < high {
		mid := low + (high-low)/2
		count := 0
		i, j := 0, n-1

		//for each round calculate how many elements are smaller than mid
		for ; i < m; i++ {
			//count at the row that has mid possibly within it
			for j >= 0 && mid < matrix[i][j] {
				j--
			}
			count += j + 1
		}
		//count < k, search in higher parts
		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
