package problem1033

import "sort"

func numMovesStones(a int, b int, c int) []int {
    xyz := []int{a,b,c}
    sort.Ints(xyz)
    if xyz[2] - xyz[0] == 2{
        return []int{0,0}
    }
    result := []int{2,xyz[2] - xyz[0] - 2}
    if xyz[1] - xyz[0] <= 2 || xyz[2] - xyz[1] <= 2{
        result[0] = 1
    }
    return result
}

