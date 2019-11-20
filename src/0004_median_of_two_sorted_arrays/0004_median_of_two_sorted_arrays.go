package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    A, B := nums1, nums2
    if len(nums1) > len(nums2){
        A,B = nums2, nums1
    }
    
    m,n := len(A), len(B)
    iMin, iMax := 0, m
    half := (m+n+1)/2
    
    var i,j int
    for (iMin<=iMax){
        i = (iMin+iMax)/2
        j = half - i
        
        if i < m && B[j-1] > A[i]{
            iMin = i + 1 // i is too small
        }else if i > 0 && A[i-1] > B[j]{
            iMax = i - 1 // i is too big
        }else{
            break
        }
    }
        var left,right int
        if i == 0{
            left = B[j-1]
        }else if j == 0{
            left = A[i-1]
        }else if A[i-1] > B[j-1]{
            left = A[i-1]
        }else{
            left = B[j-1]
        }

        if (m+n)%2 != 0{
            return float64(left)
        }

        if i == m{
            right = B[j]
        }else if j == n{
            right = A[i]
        }else if A[i] < B[j]{
            right = A[i]
        }else{
            right = B[j]
        }
        
        return float64(left+right)/2.0
}

