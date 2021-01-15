package easy1002

import (
	"math"
)

func commonChars(A []string) []string {
	dict := make(map[rune]int)
	for i, word := range A {
		dictByWord := make(map[rune]int)
		for _, r := range word {
			if _, ok := dictByWord[r]; ok {
				dictByWord[r]++
			} else {
				dictByWord[r] = 1
			}
		}
		if i == 0 {
			dict = dictByWord
		} else {
			for k, v1 := range dict {
				if v2, ok := dictByWord[k]; ok {
					dict[k] = int(math.Min(float64(v1), float64(v2)))
				} else {
					delete(dict, k)
				}
			}
		}
		dictByWord = make(map[rune]int)
	}
	var ans []string
	for k, v := range dict {
		if v > 1 {
			var arr []string
			for range make([]int, v) {
				arr = append(arr, string(k))
			}
			ans = append(ans, arr...)
		} else {
			ans = append(ans, string(k))
		}
	}
	return ans
}
