package problem0447

func numberOfBoomerangs(points [][]int) int {
    result := 0
    for i, v:= range points{
        helperMap := make(map[int]int, len(points))
        for j, k:= range points{
            if i == j{
                continue
            }
            distance := (v[0] - k[0]) * (v[0] - k[0]) + (v[1] - k[1]) * (v[1] - k[1]) 
            if count, prs := helperMap[distance]; prs{
                result += count * 2
                helperMap[distance]++
            }else{
                helperMap[distance] = 1
            }
        }
    }
    return result
}
