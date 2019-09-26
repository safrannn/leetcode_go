package problem0520

func detectCapitalUse(word string) bool {   
    count := 0
    
    for _,v := range word{
        if v >= 'A' && v <= 'Z' {
        count++
        }
    }
    
    if count == 0 || count == len(word){
        return true
    }
    if word[0] >= 'A' && word[0] < 'Z' && count == 1{
        return true
    }
    
    return false
}