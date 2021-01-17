package easy543

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := -1
	var search func(root *TreeNode) int
	search = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := search(root.Left)
		right := search(root.Right)
		ans = max(ans, (left + right + 1))
		return max(left, right) + 1
	}
	search(root)
	return ans - 1
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
