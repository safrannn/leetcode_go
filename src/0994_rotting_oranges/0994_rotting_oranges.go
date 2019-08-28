package problem0994

import "fmt"

func orangesRotting(grid [][]int) int {
	//BFS
	row, col := len(grid), len(grid[0])
	freshCount := 0
	rottingQueue := make([][]int, 0, row*col)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				rottingQueue = append(rottingQueue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}
	if freshCount == 0 {
		return 0
	}
	roundCount := 0

	for len(rottingQueue) > 0 {
		fmt.Println(rottingQueue)
		newQueue := [][]int{}
		for i := range rottingQueue {
			x, y := rottingQueue[i][0], rottingQueue[i][1]
			if x-1 >= 0 && grid[x-1][y] == 1 {
				newQueue = append(newQueue, []int{x - 1, y})
				grid[x-1][y] = 2
				freshCount--
			}
			if x+1 < row && grid[x+1][y] == 1 {
				newQueue = append(newQueue, []int{x + 1, y})
				grid[x+1][y] = 2
				freshCount--
			}

			if y-1 >= 0 && grid[x][y-1] == 1 {
				newQueue = append(newQueue, []int{x, y - 1})
				grid[x][y-1] = 2
				freshCount--
			}
			if y+1 < col && grid[x][y+1] == 1 {
				newQueue = append(newQueue, []int{x, y + 1})
				grid[x][y+1] = 2
				freshCount--
			}
		}
		rottingQueue = newQueue
		roundCount++
	}

	if freshCount > 0 {
		return -1
	}

	return roundCount - 1
}
