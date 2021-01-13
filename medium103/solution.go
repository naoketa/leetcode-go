package medium103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := queue{}
	q.enqueue(root)
	var ans [][]int
	for q.size() != 0 {
		size := q.size()
		row := []int{}
		for range make([]int, size) {
			curret := q.dequeue()
			if len(ans)%2 == 1 {
				row = append([]int{curret.Val}, row...)
			} else {
				row = append(row, curret.Val)
			}

			if curret.Left != nil {
				q.enqueue(curret.Left)

			}
			if curret.Right != nil {
				q.enqueue(curret.Right)
			}
		}
		ans = append(ans, row)
	}
	return ans
}

type queue []*TreeNode

func (q *queue) enqueue(val *TreeNode) {
	*q = append((*q), val)
}

func (q *queue) dequeue() *TreeNode {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *queue) size() int {
	return len(*q)
}
