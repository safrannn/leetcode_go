package problem1281

func subtractProductAndSum(n int) int {
    if n < 10{
        return 0
    }
    sum,product := 0 ,1
    for n > 0{
        number := n % 10
        sum += number
        product *= number
        n /= 10
    }
    
    return product - sum
}