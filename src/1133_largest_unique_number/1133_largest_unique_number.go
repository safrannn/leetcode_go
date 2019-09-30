package problem1133

func largestUniqueNumber(A []int) int {
    helper := make([]int,1001)
    for _,v := range A{
        helper[v]++
    }

    for i:= 1000; i >= 0; i--{
        if helper[i] == 1{
            return i
        }
    }
    return -1
}