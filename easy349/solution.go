package easy349

func intersection(nums1 []int, nums2 []int) []int {

	set := make(map[int]struct{})
	for _, v := range nums1 {
		set[v] = struct{}{}
	}

	ansSet := make(map[int]struct{})
	for _, v := range nums2 {
		_, ok := set[v]
		if ok {
			ansSet[v] = struct{}{}
		}
	}
	ans := []int{}
	for k := range ansSet {
		ans = append(ans, k)
	}

	return ans
}
