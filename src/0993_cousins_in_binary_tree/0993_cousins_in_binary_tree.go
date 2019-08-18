package problem0993

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var queue []*TreeNode
	queue = append(queue, root)
	return helper(x, y, queue)
}

func helper(x int, y int, queue []*TreeNode) bool {
	if len(queue) == 0 {
		return false
	}
	newQueue := []*TreeNode{}
	newMap := make(map[int]int) //key: child node value, value: parent node value
	for _, v := range queue {
		if v.Left != nil {
			newMap[v.Left.Val] = v.Val
			newQueue = append(newQueue, v.Left)
		}
		if v.Right != nil {
			newMap[v.Right.Val] = v.Val
			newQueue = append(newQueue, v.Right)
		}
	}
	vx, prsx := newMap[x]
	vy, prsy := newMap[y]
	if prsx && prsy && vx != vy {
		return true
	} else {
		return helper(x, y, newQueue)
	}
}
