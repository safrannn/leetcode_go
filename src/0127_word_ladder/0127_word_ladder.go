package problem0127

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// use Bidirectional BFS
	wordMap := make(map[string]struct{})
	for _,word := range wordList{
		wordMap[word] = struct{}{}
	}
	if _,prs := wordMap[endWord]; !prs{
		return 0
	}

	src, dst := make(map[string]struct{}),make(map[string]struct{})
	src[beginWord] = struct{}{}
	dst[endWord] = struct{}{}

	result := 1
	prs := false
	for len(src) != 0 && len(dst) != 0{
		result++
		if len(src) > len(dst){
			src, dst = dst, src 
		}

		for w := range src{
			delete(wordMap, w)
		}

		newSrc := make(map[string]struct{})
		for word := range src{
			wordByte := []byte(word)

			for i :=0; i < len(wordByte); i++{
				for j := 0; j < 26; j++{
					wordByte[i] = 'a' + byte(j)
					target := string(wordByte)

					if _, prs = dst[target]; prs{
						return result
					} else if _,prs = wordMap[target];prs{
							newSrc[target] = struct{}{}
					}
				}
				wordByte[i] = word[i]
			}
		}
		src = newSrc
	}

	return 0
}