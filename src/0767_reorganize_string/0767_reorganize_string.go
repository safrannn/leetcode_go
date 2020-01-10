package problem0767

import "sort"

func reorganizeString(S string) string {
	n := len(S)
	count := make([]int,26)
	// right two digit left for character, others digits for count
	for i:=0; i<26; i++{
		count[i] += i
	}
	for _,v := range S{
		count[int(v-'a')] += 100 
	}

	sort.Ints(count)
	if count[25] / 100 > (n+1) / 2{
		return ""
	}

	result := make([]byte, n)
	currentIndex := 1

	for _,v := range count{
		character := byte('a' + v % 100)
		characterCount := v / 100

		for i := 0; i < characterCount; i++{
			if currentIndex >= n{
				currentIndex = 0
			}
			result[currentIndex] = character
			currentIndex += 2
		}
	}

	return string(result)
}