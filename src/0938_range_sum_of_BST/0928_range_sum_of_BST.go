package problem0938

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val >= L && root.Val <= R {
			sum += root.Val
		}
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)

	return sum
}
