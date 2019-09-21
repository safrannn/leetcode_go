package problem0985

func sumEvenAfterQueries(A []int, queries [][]int) []int {
    currentEvenSum := 0
    result := make([]int, len(A))
    
    for _,v :=  range A{
        if v % 2 == 0{
            currentEvenSum += v
        }
    }
    
    for i := range queries{
        val := queries[i][0]
        index := queries[i][1]
        
        if A[index] % 2 == 0{
            currentEvenSum -= A[index]
        }
        
        A[index] += val
        
        if A[index] % 2 == 0{
            currentEvenSum += A[index]
        }
        result[i] = currentEvenSum
    }
    return result
}