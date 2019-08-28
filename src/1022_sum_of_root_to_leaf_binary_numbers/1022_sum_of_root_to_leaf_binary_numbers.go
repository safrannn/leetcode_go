package problem1022

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, n int) int {
	if root == nil {
		return 0
	}

	if root.Left == root.Right {
		return n + root.Val
	}

	n = (root.Val + n) << 1

	return helper(root.Left, n) + helper(root.Right, n)
}
