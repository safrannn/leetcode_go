package problem0009

func isPalindrome(x int) bool {
    if x < 0{
        return false
    }else if x < 10{
        return true
    }
    compare, original := 0, x
    
    for original > 0{
        compare = compare * 10 + original % 10
        original /= 10
    }
    
    return compare == x
}