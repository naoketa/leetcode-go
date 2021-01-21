package easy021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	tmp := &ListNode{}
	ans := tmp
	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val <= l2.Val) {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	return ans.Next
}
