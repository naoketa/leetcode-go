package easy206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		curr := head
		head = head.Next
		curr.Next = prev
		prev = curr
	}
	return prev
}
