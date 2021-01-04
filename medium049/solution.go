package medium049

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	keyToWord := make(map[string][]string)

	for _, word := range strs {
		key := getSortedStr(word)
		arr, _ := keyToWord[key]
		keyToWord[key] = append(arr, word)
	}

	ans := [][]string{}
	for _, words := range keyToWord {
		ans = append(ans, words)
	}
	return ans
}

func getSortedStr(word string) string {
	strs := strings.Split(word, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}
