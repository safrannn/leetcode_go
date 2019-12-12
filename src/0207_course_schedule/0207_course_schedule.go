package problem0207

func canFinish(numCourses int, prerequisites [][]int) bool {
    coursePrereq := make([][]int,numCourses)
    for _,v := range prerequisites{
        coursePrereq[v[0]] = append(coursePrereq[v[0]],v[1])
    }
    
    visited := make(map[int]int)
    
    for i := 0; i < numCourses; i++{
        if hasCycle(i, coursePrereq, visited){
            return false
        }
    }
    
    return true
}

//dfs
func hasCycle(currentCourse int, coursePrereq [][]int, visited map[int]int)bool{
    if k,prs := visited[currentCourse];prs{
        if k == -1{
            return false
        }else if k == 1{
            return true
        }
    }
    
    visited[currentCourse] = 1
    
    for _,v := range coursePrereq[currentCourse]{
        if hasCycle(v, coursePrereq, visited){
            return true
        }
    }
    visited[currentCourse] = -1
    return false
}













