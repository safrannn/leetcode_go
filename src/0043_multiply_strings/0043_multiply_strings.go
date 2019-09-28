package problem0043

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0"{
        return "0"
    }
    // do a multiplication like in junior high with paper
    result := make([]int, len(num1) + len(num2))
    for i := len(num1) - 1; i >= 0; i--{
        v1 := int(num1[i] - '0')
        for j := len(num2) - 1; j >= 0 ; j --{
            mult := v1 * int(num2[j] - '0')
            p1, p2 := i + j, i + j + 1
            // add the old result[p2] for easier carry calculaton
            sum := mult + result[p2]
            result[p1] += sum/10
            result[p2] = sum%10
        }
    }
    resultString := ""
    for _,v := range result{
        resultString += strconv.Itoa(v)
    }
    //remove leading zero
    resultString = strings.TrimLeft(resultString, "0")
    return resultString
}

