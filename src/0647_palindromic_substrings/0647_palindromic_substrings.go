package problem0647

func countSubstrings(s string) int {
	count := 0
	for i:=0; i < len(s);i++{
		count += checkPalindrome(s,i,i)
		count += checkPalindrome(s,i,i+1)
	}
	print(count)
    return count
}

func checkPalindrome(s string,i int,j int)int{
	count := 0
	for i >= 0 && j < len(s) && s[i] == s[j]{
		count++
		i--
		j++
	}
	return count
}