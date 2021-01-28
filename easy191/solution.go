package easy191

func hammingWeight(num uint32) int {
	var ans int
	for num != 0 {
		ans += int(num) & 1
		num >>= 1
	}
	return ans
}
