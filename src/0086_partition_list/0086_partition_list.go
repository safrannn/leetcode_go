package problem0086

import "leetcode_go/types"

type ListNode = types.ListNode

func partition(head *ListNode, x int) *ListNode {
    dummy1, dummy2 := &ListNode{}, &ListNode{}
    ptr1, ptr2 := dummy1, dummy2
    for ptr := head; ptr != nil; ptr = ptr.Next{
        if ptr.Val < x{
            ptr1.Next = ptr
            ptr1 = ptr
        }else{
            ptr2.Next = ptr
            ptr2 = ptr
        }
    }
    ptr1.Next = dummy2.Next
    ptr2.Next = nil
    return dummy1.Next    
}