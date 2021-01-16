package easy234

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	nodes := head
	var tmp []int
	for nodes != nil {
		tmp = append(tmp, nodes.Val)
		nodes = nodes.Next
	}
	ans := true
	fmt.Println(head.Val)
	for i := len(tmp) - 1; i > len(tmp)/2-1; i-- {
		if tmp[i] == head.Val {
			head = head.Next
			continue
		}
		return false
	}
	return ans
}
