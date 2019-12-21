package problem0518

func change(amount int, coins []int) int {
    coinCombination := make([]int, amount+1)
    coinCombination[0] = 1
    
    for _, coinValue := range coins{
        for money := coinValue; money <= amount; money++{
            coinCombination[money] += coinCombination[money - coinValue]
        }
    }
    
    return coinCombination[amount]
}