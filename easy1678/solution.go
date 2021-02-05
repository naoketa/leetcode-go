package easy1678

func interpret(command string) string {
	var ans []byte

	for i := 0; i < len(command); i++ {
		switch {
		case command[i] == 'G':
			ans = append(ans, 'G')
		case command[i+1] == 'a':
			ans = append(ans, []byte{'a', 'l'}...)
			i += 3
		default:
			ans = append(ans, 'o')
			i++
		}
	}
	return string(ans)
}
