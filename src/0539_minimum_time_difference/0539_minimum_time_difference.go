package problem0539

import (
	"strings"
	"strconv"
	"sort"
)

func findMinDifference(timePoints []string) int {
    minutes := make([]int, len(timePoints)) // make(Type, len, cap)
    for i, v := range timePoints {
        minutes[i] = timeToMinute(v)
    }
    sort.Ints(minutes)
    min := minutes[0] + 1440 - minutes[len(minutes)-1]
    for i := 1; i < len(minutes); i++ {
        cur := minutes[i] - minutes[i-1]
        if cur < min {
            min = cur
        }
    }
    return min
}
func timeToMinute(s string)int{
	hm := strings.Split(s,":")
	hours,_ := strconv.Atoi(hm[0])
	minutes,_ := strconv.Atoi(hm[1])
	return hours * 60 + minutes
}