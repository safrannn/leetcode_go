package problem0492

import "math"

func constructRectangle(area int) []int {
    mid := int(math.Sqrt(float64(area)))
    for area % mid != 0{
        mid--    
    }
    return []int{area / mid, mid}
}