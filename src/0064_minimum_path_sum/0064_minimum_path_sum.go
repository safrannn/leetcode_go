package problem0064

func minPathSum(grid [][]int) int {
    // dp: dp(i,j)=grid(i,j)+min(dp(i+1,j),dp(i,j+1))
    // starting from bottom right, traverse back and fill the matrix with required minimum sum, at every step, move either rightwards or downwards
    // -> space and time O(mn), use slice instead of matrix
    // -> time O(mn), -> space O(1) store in original matrix
    
    m, n := len(grid), len(grid[0])
    
    for i := m - 1; i >= 0; i --{
        for j := n - 1; j >= 0; j--{
            // special case: bottom cases and right most cases
            if i != m - 1 && j != n - 1{
                grid[i][j] += min(grid[i+1][j], grid[i][j+1])
            }else if i == m - 1 && j != n - 1{
                grid[i][j] += grid[i][j + 1]
            }else if i != m - 1 && j == n - 1{
                grid[i][j] += grid[i + 1][j]
            }
        }
    }
    return grid[0][0]
}

func min(a,b int)int{
    if a < b{
        return a
    }
    return b
}