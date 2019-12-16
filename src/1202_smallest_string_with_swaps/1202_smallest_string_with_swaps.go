package problem1202

import (
	"leetcode_go/types"
	"sort"
)

type UnionSet = types.UnionSet

func smallestStringWithSwaps(s string, pairs [][]int) string {
    // create a union set and union pairs
    var u UnionSet
    u.Init(len(s))
    
    for _,v := range pairs{
        u.Union(v[0],v[1])
    }
    
    // use a map to record letters for each set
    dict := make(map[int][]byte,len(s))
    for i := range s{
        parent := u.Find(i)
        dict[parent] = append(dict[parent], s[i])
    }
    
    // sort each set of letters afterwards
    for _,v := range dict {
        sort.Slice(v, func(i, j int) bool {
            return v[i] < v[j]
        })
    }
    
    // put in letters from the smallest for each set
    result := make([]byte,len(s))
    for i := range s{
        parent := u.Find(i)
        result[i] = dict[parent][0]
        dict[parent] = dict[parent][1:]
    }
    return string(result)
}

