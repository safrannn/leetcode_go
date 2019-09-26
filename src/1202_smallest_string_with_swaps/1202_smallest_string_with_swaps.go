package problem1202

func smallestStringWithSwaps(s string, pairs [][]int) string {
    u := make(unionSet, len(s))
    u.init()
    
    // connect all pairs
    for _,v := range pairs{
        u.union(v[0],v[1])
    }
    
    // put characters with the same group into map
    helperMap := make(map[int][]byte)
    for i := range s{
        helperMap[u.find(i)] = append(helperMap[u.find(i)], s[i])
    }
    
    // sort each group
    for i := range helperMap{
        sort.Slice(helperMap[i], func(k int, j int) bool { return helperMap[i][k] > helperMap[i][j] })
    }
    
    //generate result by adding sorted character whose index belongs to the same group
    result := []byte{}
    for i := range s{
        current := helperMap[u.find(i)]
        result = append(result, current[len(current) - 1])
        helperMap[u.find(i)] = helperMap[u.find(i)][:len(current)-1]
    }
    return string(result)
}

type unionSet []int

func (u unionSet) init(){
    for i:=0; i<len(u); i++{
        u[i] = i
    }
}

func (u unionSet) find(x int)int{
    if u[x] == x{
        return x
    }
    return u.find(u[x])
}

func (u unionSet) union(x, y int){
    u[u.find(x)] = u.find(y)
}