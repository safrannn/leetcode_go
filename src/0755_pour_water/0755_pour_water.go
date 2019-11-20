package problem0755

func pourWater(heights []int, V int, K int) []int {
    for i := 0; i < V; i++{
        current := K
        
        for current > 0 && heights[current] >= heights[current-1]{
            current--
        }
        
        
        for current < len(heights) - 1 && heights[current] >= heights[current+1]{
            current++
        }
        
        for current > K && heights[current] >= heights[current-1]{
            current--
        }
        
        heights[current]++
    }
    return heights
}