package problem0965

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Left.Val != root.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val != root.Val {
			return false
		}
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
