package easy559

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	q := queue{}
	q.enqueue(root)

	var depth int
	for q.size() != 0 {
		size := q.size()
		for range make([]int, size) {
			head := q.dequeue()
			for _, child := range head.Children {
				q.enqueue(child)
			}
		}
		depth++
	}

	return depth
}

type queue []*Node

func (q *queue) enqueue(val *Node) {
	*q = append((*q), val)
}

func (q *queue) dequeue() *Node {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *queue) size() int {
	return len(*q)
}
