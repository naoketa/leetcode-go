package easy001

func twoSum(nums []int, target int) []int {
	numToIdx := make(map[int]int)

	for idx, v := range nums {
		pair := target - v
		pairIdx, ok := numToIdx[pair]
		if ok {
			return []int{pairIdx, idx}
		}
		numToIdx[v] = idx
	}
	return []int{}
}
