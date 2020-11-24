package problem0337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func rob(root *TreeNode) int {
    x,y := traverse(root);
    return max(x,y);
}

func traverse(root *TreeNode)(int,int){
    if root == nil{
        return 0,0;
    }
    left_x,left_y := traverse(root.Left);
    right_x,right_y := traverse(root.Right);
    return root.Val + left_y + right_y, max(left_x,left_y) + max(right_x,right_y);
}

func max(a,b int)int{
    if a > b{
        return a
    }else{
        return b
    }
}
