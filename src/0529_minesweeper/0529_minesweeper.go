package problem0529

func updateBoard(board [][]byte, click []int) [][]byte {
    if board[click[0]][click[1]] == 'M'{
        board[click[0]][click[1]] = 'X'
        return board
    }
    dfs(board, click[0],click[1])
    return board
}

func dfs(board [][]byte, i int, j int){
    n, m := len(board), len(board[0])
    if i < 0 || i >= n || j < 0 || j >= m ||board[i][j] != 'E'{
        return
    }
    
    mineCount := checkSurroundings(board, i, j)
    if checkSurroundings(board, i, j) > 0{
        board[i][j] = byte(mineCount) + '0'
    }else{
        board[i][j] = 'B'
        dfs(board,i-1,j-1)
        dfs(board,i,j-1)
        dfs(board,i+1,j-1)
        dfs(board,i-1,j)
        dfs(board,i+1,j)
        dfs(board,i-1,j+1)
        dfs(board,i,j+1)
        dfs(board,i+1,j+1)
    }   
}

func checkSurroundings(board [][]byte, r int, c int) int{
    mineCount := 0
    for i := r - 1; i <= r + 1; i++ {
        for j := c - 1; j <= c + 1; j++ {
            if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
                continue
            }
            if (board[i][j] == 'M' || board[i][j] == 'X') {
                mineCount++
            }
        }
    }
    return mineCount
}

