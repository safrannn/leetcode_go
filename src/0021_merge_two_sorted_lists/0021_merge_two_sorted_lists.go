package problem0021

import (
	"leetcode_go/types"
)

// ListNode is defined in leetcode_go/types
type ListNode = types.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }else if l2 == nil{
        return l1
    }
    
    if l1.Val < l2.Val{
        l1.Next = mergeTwoLists(l1.Next,l2)
        return l1
    }else{
        l2.Next = mergeTwoLists(l1,l2.Next)
        return l2
    }
}