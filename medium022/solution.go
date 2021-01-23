package medium022

func generateParenthesis(n int) []string {
	var ans []string
	var rec func(left, right int, str string)
	rec = func(left, right int, str string) {
		if right > left || left > n || right > n {
			return
		}
		if left == n && right == n {
			ans = append(ans, str)
		}
		rec(left+1, right, str+"(")
		rec(left, right+1, str+")")
	}
	rec(0, 0, "")
	return ans
}
