package problem1031

func maxSumTwoNoOverlap(A []int, L int, M int) int {
    for i :=1; i < len(A);i++{
        A[i] += A[i-1]
    }
    
    result := A[L+M-1]
    Lmax := A[L-1]
    Mmax := A[M-1]
    for i := L+M; i < len(A);i++{
        //Lmax, max sum of contiguous L elements before the last M elements.
        Lmax = max(Lmax, A[i-M] - A[i-L-M])
        //Mmax, max sum of contiguous M elements before the last L elements/
        Mmax = max(Mmax, A[i-L] - A[i-L-M])
        
        result = max(result,max(Lmax + A[i] - A[i - M], Mmax + A[i] - A[i - L]) )
    }
    return result
}

func max(a,b int)int{
    if a > b {
        return a
    }
    return b
}