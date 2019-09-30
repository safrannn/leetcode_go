package problem1118

import ("strconv"
    "time")
    
func numberOfDays(Y int, M int) int {
    year := strconv.Itoa(Y)
    month := strconv.Itoa(M)
    if M < 10{
        month = "0" + month
    }
    current,_ := time.Parse("2006-01-02", year+"-"+month+"-01")
    var nextYear string
    var nextMonth string
    if M == 12{
        nextYear = strconv.Itoa(Y + 1)
        nextMonth = "01"
    }else{
        nextYear = strconv.Itoa(Y)
        nextMonth = strconv.Itoa(M+1)
        if M+1< 10{
            nextMonth = "0" + nextMonth
        }
    }
    next,_ := time.Parse("2006-01-02", nextYear+"-"+nextMonth+"-01")
    return int(next.Sub(current).Hours())/24
}