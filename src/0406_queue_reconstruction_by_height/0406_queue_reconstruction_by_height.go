package problem0406

import ("sort")

func reconstructQueue(people [][]int) [][]int {
    result := [][]int{}
    sort.Slice(people, func(i, j int) bool{
        if people[i][0] == people[j][0]{
            return people[i][1] < people[j][1]
        }
        return people[i][0] > people[j][0]
    })
    
    for _,v := range people{
        result = append(result,[]int{})
        copy(result[v[1]+1:], result[v[1]:])
        result[v[1]] = v
    }
    
    return result
}