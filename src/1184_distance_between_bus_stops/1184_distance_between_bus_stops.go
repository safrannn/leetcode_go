package problem1184

func distanceBetweenBusStops(distance []int, start int, destination int) int {
    if start > destination{
        destination, start = start, destination
    }
    inbetween,rest := 0,0
    for i := 0; i < len(distance); i++{
        if i >= start && i < destination{
            inbetween += distance[i]
        }else{
            rest += distance[i]
        }    
    }
    if inbetween > rest{
        return rest
    }
    return inbetween
}