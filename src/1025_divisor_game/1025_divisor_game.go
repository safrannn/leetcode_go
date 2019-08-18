package problem1025

func divisorGame(N int) bool {
	//if alice loose for N, alice win for N+1
	//alice win for 2, loose for 3
	return N%2 == 0
}
