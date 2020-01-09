package problem0498

func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0{
        return []int{}
    }
    row,col := 0,0
    m,n := len(matrix), len(matrix[0])
    result := make([]int, m * n)
    
    for i :=0; i<len(result);i++{
        result[i] = matrix[row][col]
        if (row+col)%2 == 0{
            if col == n - 1{
                row++
            }else if row == 0{
                col++
            }else{
                col++
                row--
            }
        }else {
            if row == m - 1{
                col++
            }else if col == 0{
                row++
            }else{
                col--
                row++
            }
        }
    }
    return result
}
