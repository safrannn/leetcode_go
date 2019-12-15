package problem0210

func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int,numCourses)
    for _,v := range prerequisites{
        graph[v[0]] =  append(graph[v[0]], v[1])
    }
    // 0: unknown, 1: visited, -1: visiting
    visit := make([]int,numCourses)
    answer := []int{}
    for i :=0; i< numCourses; i++{
        if dfs(i,&graph,&visit,&answer){
            return []int{}
        }
    }
    return answer
}

// true for cycle
func dfs(current int, graph *[][]int, visit *[]int, answer *[]int)bool {
    if (*visit)[current] == 1{
        return false
    }
    if (*visit)[current] == -1{
        return true
    }
    
    (*visit)[current] = -1
    for _,v := range (*graph)[current]{
        if dfs(v, graph, visit, answer){
            return true
        }
    }
    
    (*visit)[current] = 1
    *answer = append(*answer, current)
    return false    
}