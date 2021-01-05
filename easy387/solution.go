package easy387

func firstUniqChar(s string) int {
	strs := []byte(s)
	m := make(map[byte]int)
	for i, b := range strs {
		_, ok := m[b]
		if ok {
			m[b] = -1
			continue
		}
		m[b] = i
	}
	min := len(s)
	for _, v := range m {
		if v != -1 && min > v {
			min = v
		}
	}

	if min == len(s) {
		return -1
	}
	return min
}
