package problem0007

func reverse(x int) int {
    result := 0 
    sign := true
    if x < 0 {
        x = -1 * x
        sign = false
    }
    for x > 0{
        remainder := x%10
        result = result * 10 + remainder
        x /= 10
    }
    if result > 2147483647 ||result < -2147483648{
        return 0
    }else if sign == false{
        return result * -1
    }else{
        return result
	}
}