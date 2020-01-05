package problem1309

func freqAlphabets(s string) string {
    result := []byte{}
    for i := 0; i < len(s);i++{
        if i < len(s) - 2 && s[i+2] == '#'{
            result = append(result, 'j' + (s[i] - '1') * 10 + s[i+1] -'0')
            i += 2
        }else{
            result = append(result, 'a' + (s[i] - '1'))
        }
    }
    return string(result)
}