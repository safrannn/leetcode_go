package problem1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
    satisfy := 0
    extra := 0
    extraMax := 0
    for i := 0; i < len(customers); i++{
        if grumpy[i] == 0{
            satisfy += customers[i]
        }else{
            extra += customers[i]
        }
        
        if i >= X{
            extra -= grumpy[i-X] * customers[i-X]
        }
        if extra > extraMax{
            extraMax = extra
        }
    }
    return satisfy + extraMax
}