package medium098

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTRec(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTRec(root *TreeNode, lower, higher int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= higher {
		return false
	}
	return isValidBSTRec(root.Left, lower, root.Val) && isValidBSTRec(root.Right, root.Val, higher)

}
