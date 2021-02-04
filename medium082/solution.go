package medium082

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummpy := ans
	var pre *ListNode
	current := head
	for current != nil {
		case1 := pre != nil && current.Next != nil && current.Val != pre.Val && current.Val != current.Next.Val
		case2 := pre == nil && current.Next != nil && current.Val != current.Next.Val
		case3 := current.Next == nil && current.Val != pre.Val
		if case1 || case2 || case3 {
			dummpy.Next = current
			dummpy = dummpy.Next
		}
		pre = current
		current = current.Next
	}
	dummpy.Next = nil
	return ans.Next
}
