package problem1185

import "time"

func dayOfTheWeek(day int, month int, year int) string {
    timeParse := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return timeParse.Weekday().String()
}