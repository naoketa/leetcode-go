package medium153

import "math"

func findMin(nums []int) int {
	min := math.MaxInt64
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}
