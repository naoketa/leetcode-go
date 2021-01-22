package easy053

import "math"

func maxSubArray(nums []int) int {
	max, sum := math.MinInt64, 0
	for _, v := range nums {
		sum += v
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
