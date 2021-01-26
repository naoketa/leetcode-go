package medium739

func dailyTemperatures(T []int) []int {
	stack := &stack{}
	ans := make([]int, len(T))
	for i, v := range T {
		for stack.len() != 0 && T[stack.last()] < v {
			last := stack.pop()
			ans[last] = i - last
		}
		stack.push(i)
	}
	return ans
}

type stack []int

func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) last() int {
	return (*s)[len(*s)-1]
}

func (s *stack) len() int {
	return len(*s)
}
