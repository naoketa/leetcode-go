package easy387

func firstUniqChar(s string) int {
	runeToCnt := make(map[rune]int)
	for _, v := range s {
		runeToCnt[v]++
	}
	for i, v := range s {
		if runeToCnt[v] == 1 {
			return i
		}
	}
	return -1
}
