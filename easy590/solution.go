package easy590

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	s := stack{}
	s.push(root)

	var ans []int
	for s.len() > 0 {
		f := s.pop()
		ans = append([]int{f.Val}, ans...)
		for _, v := range f.Children {
			s.push(v)
		}
	}
	return ans
}

type stack []*Node

func (s *stack) push(v *Node) {
	*s = append(*s, v)
}

func (s *stack) pop() *Node {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) len() int {
	return len(*s)
}
