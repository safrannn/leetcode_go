package problem1119

func removeVowels(S string) string {
    sByte := []byte(S)
    newByte := []byte{}
    for _,v := range sByte{
        if v != 'a' && v != 'e' && v != 'i' && v != 'o' && v != 'u' {
            newByte = append(newByte,v)
        }
    }
    return string(newByte)
}