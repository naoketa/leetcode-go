package medium560

func subarraySum(nums []int, k int) int {
	preSum := make(map[int]int)
	preSum[0] = 1

	sum, ans := 0, 0
	for _, v := range nums {
		sum += v
		if ps, ok := preSum[sum-k]; ok {
			ans += ps
		}
		preSum[sum]++
	}
	return ans
}
