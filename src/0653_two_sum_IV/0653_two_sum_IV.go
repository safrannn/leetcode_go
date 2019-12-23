package problem0653

func findTarget(root *TreeNode, k int) bool {
    numberMap := make(map[int]bool)
    return traverse(&numberMap,root,k) 
}

func traverse(numberMap *map[int]bool,root *TreeNode,k int)bool{
    if root == nil{
        return false
    }
    if (*numberMap)[root.Val]{
        return true
    }
    (*numberMap)[k-root.Val] = true
    return traverse(numberMap,root.Left,k) || traverse(numberMap,root.Right,k)
}


// func findTarget(root *TreeNode, k int) bool {
//     array := []int{}
//     treeToArray(root, &array)
    
//     left, right := 0, len(array)-1
//     for left < right{
//         sum := array[left] + array[right]
//         if sum < k{
//             left++
//         }else if sum > k{
//             right--
//         }else{
//             return true
//         }
//     }
//     return false
// }

// func treeToArray(root *TreeNode, array *[]int){
//     if root == nil{
//         return
//     }
//     treeToArray(root.Left, array)
//     *array = append(*array,root.Val)
//     treeToArray(root.Right, array)
// }