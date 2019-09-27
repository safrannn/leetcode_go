package problem1189

import ("sort")

func maxNumberOfBalloons(text string) int {
    // need a * 1, b * 1, l * 2, o * 2, n * 1
    balloonCharCount := make([]int,5)
    for _, v := range text{
        switch rune(v) {
            case 'a':
            balloonCharCount[0]++
            case 'b':
            balloonCharCount[1]++
            case 'l':
            balloonCharCount[2]++
            case 'o':
            balloonCharCount[3]++
            case 'n':
            balloonCharCount[4]++
        }
	}
	
    // l, o appeared twice
    balloonCharCount[2] /=2
    balloonCharCount[3] /=2
    
    sort.Ints(balloonCharCount)
    return balloonCharCount[0]
}