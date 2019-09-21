package problem0754

func reachNumber(target int) int {
    // go over the goal
    // go over by even number, choose one of the step to go left
    // go over by odd number, keep going until go over by even number
    if target < 0{
        target = -target
    } 
    result := 1; sum := 0
    for (sum < target || (sum - target) % 2 == 1) {
        sum += result;
        result++;
    }
    return result - 1;
}