package easy136

func singleNumber(nums []int) int {
	numToCnt := make(map[int]int)
	for _, v := range nums {
		if _, ok := numToCnt[v]; ok {
			numToCnt[v]++
		} else {
			numToCnt[v] = 1
		}
	}
	for k, v := range numToCnt {
		if v < 2 {
			return k
		}
	}
	return -1
}
