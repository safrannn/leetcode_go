package problem0328

import (
	"leetcode_go/types"
)

// ListNode is defined in leetcode_go/types
type ListNode = types.ListNode


func oddEvenList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    
    odd := head
    even := head.Next
    evenDummy := &ListNode{}
    evenDummy.Next = even
    
    for even != nil && even.Next != nil{
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next        
    }
    
    odd.Next = evenDummy.Next
    
    return head
}