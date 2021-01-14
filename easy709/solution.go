package easy709

func toLowerCase(str string) string {
	var ans []rune
	for _, v := range str {
		if v < 'a' && 'A' <= v {
			ans = append(ans, v+'a'-'A')
		} else {
			ans = append(ans, v)
		}
	}
	return string(ans)
}
