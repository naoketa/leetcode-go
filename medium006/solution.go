package medium006

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows < 1 {
		return ""
	}
	if numRows == 1 {
		return s
	}

	tmp := make([][]byte, numRows+1)
	step, row := 1, 0
	for i := 0; i < len(s); i++ {
		tmp[row] = append(tmp[row], s[i])
		if (row == 0 && step == -1) || (row == numRows-1 && step == 1) {
			step *= -1
		}
		row += step
	}
	var ans []byte
	for _, v := range tmp {
		ans = append(ans, v...)
	}
	return string(ans)
}
