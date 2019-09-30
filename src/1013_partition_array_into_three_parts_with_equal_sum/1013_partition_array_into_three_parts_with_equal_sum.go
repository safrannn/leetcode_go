package problem1013

func canThreePartsEqualSum(A []int) bool {
    sum :=0
    for _,v := range A{
        sum +=v
    }
    if sum % 3 != 0 {
        return false
    }
    sum /=3
    currentSum,currentTry := 0, 0
    for _,v := range A{
        currentSum += v
        if currentSum == sum {
            currentTry++
            currentSum = 0
        }
    }
    return currentTry == 3
}