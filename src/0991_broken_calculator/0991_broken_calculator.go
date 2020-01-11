package problem0991

func brokenCalc(X int, Y int) int {
    result := 0
    
    for Y > X{
        result++
        if Y % 2 == 0{
            Y /= 2
        }else{
            Y++
        }
    }
    
    return result + X - Y
}