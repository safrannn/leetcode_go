package problem1165

func calculateTime(keyboard string, word string) int {
    keyMap := make([]int, 26)
    for i,v := range keyboard{
        keyMap[v - 'a'] = i
    } 
    
    result := 0
    previous := 0
    for _,v := range word{
        result += abs(previous - keyMap[v - 'a'])
        previous = keyMap[v -'a']
    }
    return result
}

func abs(a int)int{
    if a < 0{
        return -a
    }
    return a
}