package medium003

func lengthOfLongestSubstring(s string) int {

	window := make(map[byte]struct{})
	i, j, max := 0, 0, 0
	for j < len(s) && i < len(s) {
		_, ok := window[s[j]]
		if !ok {
			window[s[j]] = struct{}{}
			if (j - i + 1) > max {
				max = j - i + 1
			}
			j++
		} else {
			delete(window, s[i])
			i++
		}
	}

	return max
}
