package easy771

func numJewelsInStones(jewels string, stones string) int {
	set := make(map[rune]struct{})
	for _, v := range jewels {
		set[v] = struct{}{}
	}

	var cnt int
	for _, v := range stones {
		if _, ok := set[v]; ok {
			cnt++
		}
	}
	return cnt
}
