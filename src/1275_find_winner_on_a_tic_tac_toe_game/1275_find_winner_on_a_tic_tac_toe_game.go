package problem1275

func tictactoe(moves [][]int) string {
    wins := [][]int{[]int{0,1,2},[]int{3,4,5},[]int{6,7,8},[]int{0,3,6}, []int{1,4,7}, []int{2,5,8}, []int{0,4,8},[]int{2,4,6}}
    a_take := make([]bool, 9)
    b_take := make([]bool, 9)
    for i,v := range moves{
        if i % 2 == 0{
            a_take[v[0] * 3 + v[1]] = true;
        }else{
            b_take[v[0] * 3 + v[1]] = true;
        }
    }
    
    for _,v:= range wins{
        if a_take[v[0]] &&a_take[v[1]]&&a_take[v[2]]{
            return "A"
        }else if b_take[v[0]] && b_take[v[1]] && b_take[v[2]]{
            return "B"
        }
    }
    if len(moves) < 9{
        return "Pending"
    }
    return "Draw"
}