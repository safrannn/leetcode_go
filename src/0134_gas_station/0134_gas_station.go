package problem0134

func canCompleteCircuit(gas []int, cost []int) int {
    totalTank, currentTank := 0,0
    startStation := 0
    for i := range gas{
        totalTank += gas[i] - cost[i]
        currentTank += gas[i] - cost[i]
        if currentTank < 0{
            startStation = i+1
            currentTank = 0
        }
    }
    if totalTank < 0{
        return -1
    }else{
        return startStation
    }
}