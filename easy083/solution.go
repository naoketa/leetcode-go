package easy083

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 1 2

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head

}
