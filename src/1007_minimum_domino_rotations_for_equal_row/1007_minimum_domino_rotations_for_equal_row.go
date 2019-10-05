problem package1007

func minDominoRotations(A []int, B []int) int {
	// first find all possible numbers of exchange
	// if count[i] >= len(A), then it is possible to form an answer
	count := count(A,B)
	result := -1
	// for all potential values
	// calculate minimum cost to align at A or B, return the min value
	for i := range count{
		if count[i] == len(A){
			cost := cost(A, B, i)
			if result == -1 || cost < result{
				result = cost
			}
		}
	}
	return result
}

func count(A []int,B []int)[]int{
	count := make([]int,7)
	for i := range A{
		count[A[i]]++
		if A[i]!=B[i]{
			count[B[i]]++
		}
	}
	return count
}

func cost(A []int, B []int, value int)int{
	costA, costB := 0, 0
	for i := range A{
		if A[i] != value{
			costA++
		}
		if B[i] != value{
			costB++
		}
	}
	if costA < costB{
		return costA
	}
	return costB
}