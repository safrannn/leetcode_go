
func highFive(items [][]int) [][]int {
    scoreMap := make(map[int][]int)
    for _,v := range items{
        id := v[0]
        score := v[1]
        if len(scoreMap[id]) < 5{
            scoreMap[id] = append(scoreMap[id],score)
        }else{
            if scoreMap[id][0] < score{
                scoreMap[id][0] = score
            }
        }
        sort.Ints(scoreMap[id])
    }

    answer := [][]int{}
    for i,v := range scoreMap{
        avg := (v[0] + v[1] + v[2] + v[3] + v[4])/5
        answer = append(answer, []int{i,avg})
    }

    sort.SliceStable(answer, func(i, j int) bool { return answer[i][0] <  answer[j][0] })
    
    return answer
}
