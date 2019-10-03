package problem0166

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
    flag := ""
	switch {
	case numerator < 0 && denominator < 0:
		numerator, denominator = -numerator, -denominator
	case numerator > 0 && denominator < 0:
		denominator, flag = -denominator, "-"
	case numerator < 0 && denominator > 0:
		numerator, flag = -numerator, "-"
	}
    
    intg, remainder := numerator / denominator, numerator % denominator
    result := flag + strconv.Itoa(intg)
    if remainder == 0 {
        return result
    }
    result += "."
    
    currentIndex := len(result) - 1
    helperMap := make(map[int]int) 
    // key: remainder * 10, value:index, if remainder start to loop, so does fractional part
    for remainder > 0{
        numerator = 10 * remainder
        intg, remainder = numerator / denominator, numerator % denominator
        
        if index, prs := helperMap[numerator]; prs{
            resultB := []byte(result + ") ")
			copy(resultB[index+1:], resultB[index:])
			resultB[index] = '('
            result = string(resultB)
			break
        }
        
        result += strconv.Itoa(intg)
        currentIndex++
        helperMap[numerator] = currentIndex
    }
    return result
}
func abs(a int)int{
    if a < 0{
        return - a
    }
    return a
}