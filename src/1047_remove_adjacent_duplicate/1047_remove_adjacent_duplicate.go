package problem1047

func removeDuplicates(S string) string {
	result := []byte{}
	for i := 0; i < len(S); i++ {
		if len(result) != 0 && result[len(result)-1] == S[i] {
			result = result[:len(result)-1]
		} else {
			result = append(result, S[i])
		}
	}
	return string(result)
}
