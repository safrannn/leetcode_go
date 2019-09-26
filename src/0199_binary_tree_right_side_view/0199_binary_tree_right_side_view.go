package problem0199

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rightSideView(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    
    result := []int{}
    stack := []*TreeNode{root}
    

    for len(stack) > 0{
        size := len(stack)
        //add value of the last node to result
        result = append(result, stack[size-1].Val)

        for i:= 0; i < size; i++{
            current := stack[0] 
            stack = stack[1:]
            if current.Left != nil{
                stack = append(stack, current.Left)
            }
            if current.Right != nil{
                stack = append(stack, current.Right)
            }
        }
    }
    return result
}