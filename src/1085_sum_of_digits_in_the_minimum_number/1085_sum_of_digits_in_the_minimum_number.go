package problem1085

func sumOfDigits(A []int) int {
    minElement := A[0]
    for i:=1; i < len(A);i++{
        if A[i] < minElement{
            minElement = A[i]
        }
    }
    
    // sum up all digits
    digitSum := 0
    for minElement > 0{
        digitSum += minElement % 10
        minElement /= 10
    }
    if digitSum % 2 == 1{
        return 0
    }
    return 1
}