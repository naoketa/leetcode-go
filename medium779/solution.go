package medium779

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}

	// even
	if K%2 == 0 {
		if kthGrammar(N-1, K/2) == 0 {
			return 1
		}
		return 0

	}

	// odd
	return kthGrammar(N-1, (K+1)/2)
}
