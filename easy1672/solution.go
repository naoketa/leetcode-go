package easy1672

func maximumWealth(accounts [][]int) int {
	max := -1
	for _, wealth := range accounts {
		var sum int
		for _, v := range wealth {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
