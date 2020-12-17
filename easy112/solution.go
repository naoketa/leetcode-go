package easy112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	onLeft := hasPathSum(root.Left, sum-root.Val)
	onRight := hasPathSum(root.Right, sum-root.Val)

	return onLeft || onRight

}
