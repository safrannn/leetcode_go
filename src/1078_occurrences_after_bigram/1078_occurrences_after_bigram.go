package problem1078

import "strings"

func findOcurrences(text string, first string, second string) []string {
    result := []string{}
    textSlice := strings.Split(text," ")
    for i:=0; i < len(textSlice)-2; i++{
        if textSlice[i] == first && textSlice[i+1] == second{
            result = append(result, textSlice[i+2])
        }
    }
    return result
}