package problem1071

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
    prefix := str1
    if len(str1) > len(str2){
        prefix = str2
    }
    
    for prefix != "" {
        if strings.HasPrefix(str1, prefix) && strings.HasPrefix(str2, prefix) && len(str1) % len(prefix) == 0 && len(str2) % len(prefix) == 0{
        return prefix
        }
            prefix = prefix[:len(prefix) - 1]

    }
    
    return ""
}

