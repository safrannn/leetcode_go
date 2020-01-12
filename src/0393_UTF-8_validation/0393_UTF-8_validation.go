package problem0393

func validUtf8(data []int) bool {
    count := 0
    
    for _,v := range data{
        if count == 0{
            if v >> 3 == 30{
                count += 3
                continue
            }
            if v >> 4 == 14{
                count += 2
                continue
            }
            if v >> 5 == 6{
                count += 1
                continue
            }
            if v >> 7 == 0{
                continue
            }
            return false
        }else{
            if v >> 6 == 2{
                count--
                continue
            }
            return false
        }
    }
    
    return count == 0
}