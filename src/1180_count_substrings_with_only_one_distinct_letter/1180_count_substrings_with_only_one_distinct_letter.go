package problem1180

func countLetters(S string) int {
    result, repeat := 0, 1
    for i:=1; i<len(S);i++{
        if S[i] != S[i-1]{
            result += repeat * (repeat+1)/2
            repeat = 0
        }
        repeat++
    }
    return result + repeat * (repeat+1)/2
}