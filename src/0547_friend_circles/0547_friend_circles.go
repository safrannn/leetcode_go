package problem0547

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

func (u unionSet) union(x,y int){
    u[x] = y
}

func findCircleNum(M [][]int) int {
    if len(M) == 0{
        return 0
    }
    
    u := make(unionSet, len(M) + 1)
    u.init()
    
    for i:= 0; i < len(M); i++{
        for j:=0; j <= i;j++{
            if M[i][j] == 1{
                a,b := u.find(i), u.find(j)
                if a != b{
                    u.union(a,b)
                }
            }
        }
    }
    
    count := 0
    
    for i:= 0; i< len(M); i++{
        if u[i] == i{
            count++
        }
    }
    
    return count
}

