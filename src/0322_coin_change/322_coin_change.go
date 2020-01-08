package problem0322

func coinChange(coins []int, amount int) int {
    coinNumbers := make([]int,amount+1)
    coinNumbers[0] = 0
    for money:=1; money<=amount; money++{
        coinNumbers[money] = amount + 1
        for _,coinValue := range coins{
            if coinValue <= money && coinNumbers[money-coinValue] + 1 < coinNumbers[money]{
                coinNumbers[money] = coinNumbers[money-coinValue] + 1 
            }
        }
    }
    if coinNumbers[amount] <= amount{
        return coinNumbers[amount]
    }else{
        return -1
    }
}