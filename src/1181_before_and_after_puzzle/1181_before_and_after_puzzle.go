package problem1181

import("strings"
        "sort"
)
	
func beforeAndAfterPuzzles(phrases []string) []string {
    result := []string{}
    
    startIndex := make(map[string][]int)
    for i,v := range phrases{
        if strings.Contains(v," "){
            startIndex[v[:strings.Index(v," ")]] = append(startIndex[v[:strings.Index(v," ")]],i)
        }else{
            startIndex[v] = append(startIndex[v],i)
        }
    }
    

    for i,v := range phrases{
        lastWord := v[strings.LastIndex(v," ")+1:]
        for _,w := range startIndex[lastWord]{
            if w != i{
                result = append(result, combine(v,phrases[w]))
            }
        }
    }
    sort.Strings(result)
    result = removeDuplicate(result)
    
    return result
}

func combine(s1,s2 string)string{
    s1Single, s2Single := !strings.Contains(s1," "), !strings.Contains(s2," ")
    if s1Single && s2Single {
        return s1
    }else if s1Single && !s2Single{
        return s2
    }else if !s1Single && s2Single{
        return s1
    }
    return s1 + s2[strings.Index(s2, " "):]
}

func removeDuplicate(s []string)[]string{
    if len(s) < 2{
        return s
    }
    s2 := []string{s[0]}
    for i :=1; i < len(s);i++{
        if s2[len(s2)-1] != s[i]{
            s2 = append(s2,s[i])
        }
    }
    return s2
}
