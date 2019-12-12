package problem1266

func minTimeToVisitAllPoints(points [][]int) int {
    result := 0
    for i:=0; i<len(points)-1;i++{
        result += max(abs(points[i][0] - points[i+1][0]), abs(points[i][1] - points[i+1][1]))
    }
    return result
}

func abs(a int)int{
    if a < 0{
        return -a
    }
    return a 
}
func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}