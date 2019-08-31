package problem0787

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	//dynamic programming https://www.youtube.com/watch?v=PLY-lbcxEjg
	//costs[src] = 0
	//cost to current destination = min(cost to current destination, cost to previous destinatoin + cost from previous to current destination)
	//use a slice to store the costs
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = math.MaxInt32
	}
	costs[src] = 0

	for i := 0; i <= K; i++ {
		currentCosts := make([]int, n)
		copy(currentCosts, costs)

		for _, v := range flights {
			currentCosts[v[1]] = getMin(currentCosts[v[1]], costs[v[0]]+v[2])
		}

		costs = currentCosts
		fmt.Println(costs)
	}

	if costs[dst] >= math.MaxInt32 {
		return -1
	}
	return costs[dst]
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
