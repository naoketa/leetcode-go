package easy020

import (
	"fmt"
)

func isValid(s string) bool {
	st := &Stack{}
	for _, v := range s {
		if st.IsEmpty() {
			st.Push(v)
			continue
		}

		switch v {
		case []rune(")")[0]:
			if last, _ := st.Pop(); last != []rune("(")[0] {
				st.Push(last)
				st.Push(v)
			}
		case []rune("}")[0]:
			if last, _ := st.Pop(); last != []rune("{")[0] {
				st.Push(last)
				st.Push(v)
			}
		case []rune("]")[0]:
			if last, _ := st.Pop(); last != []rune("[")[0] {
				st.Push(last)
				st.Push(v)
			}
		default:
			st.Push(v)
		}

	}
	return st.IsEmpty()
}

type Stack []rune

func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (rune, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
