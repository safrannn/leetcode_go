package problem0071

import "strings"

func simplifyPath(path string) string {
    paths := strings.Split(path,"/")
    stack := []string{}

    for _,v := range paths{
        if v == "." || v == ""{
            continue
        }else if v == ".."{
            if len(stack) > 0{
                stack = stack[:len(stack)-1]
            }
        }else{
            stack = append(stack,v)
        }
    }

    result := ""
    for _,v := range stack{
        result = result + "/" + v
    }
    
    if len(result) == 0{
        return "/"
    }
    return result
}