package problem0005

func longestPalindrome(s string) string {
    if len(s)<2{
        return s
    }
    result := string(s[0])
	for i := 0; i < len(s); i++ {
			result = checkPalindrome(s, i, i, result)
	}
    for i := 0; i < len(s) - 1 ; i++ {
			result = checkPalindrome(s, i, i + 1, result)
	}
	return result
}

func checkPalindrome(s string, startLeft int, startRight int, result string) string {
    subString := ""
	for startLeft >= 0 && startRight < len(s) && s[startLeft] == s[startRight]{

        subString = s[startLeft : startRight+1]
		startLeft--
		startRight++
	}
	if len(subString) > len(result) {
		result = subString
	}
	return result
}
