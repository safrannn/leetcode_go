package types

// UnionSet is disjoint set 
type UnionSet struct {
	Parent []int
}

func (u *UnionSet) Init(length int){
	u.Parent = make([]int, length)
	for k := range u.Parent{
		u.Parent[k] = k
	}
}

func (u *UnionSet) Find(x int)int{
	if u.Parent[x] != x{
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

func (u *UnionSet) Union(x, y int){
	u.Parent[u.Find(x)] = u.Find(y)
}

