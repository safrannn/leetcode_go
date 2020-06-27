package problem1323

func maximum69Number (num int) int {
    if num > 6000 && num < 9000 {
        return num + 3000
    }
    if num % 1000 > 600 && num % 1000 < 900 {
        return num + 300
    }
    if num % 100 > 60 && num % 100 < 90 {
        return num + 30
    }
    if num % 10 == 6 {
        return num + 3
    }
    return num
}