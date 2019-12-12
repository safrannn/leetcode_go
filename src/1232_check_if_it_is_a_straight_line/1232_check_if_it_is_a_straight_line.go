package problem1232

func checkStraightLine(coordinates [][]int) bool {
    c0x := coordinates[0][0]
    c0y := coordinates[0][1]
    c1x := coordinates[1][0]
    c1y := coordinates[1][1]
    
    for i:= 1; i < len(coordinates); i++{
        if (coordinates[i][0] - c0x) * (coordinates[i][1] - c1y) != (coordinates[i][0] - c1x) * (coordinates[i][1] - c0y){
            return false
        }
    } 
    return true
}