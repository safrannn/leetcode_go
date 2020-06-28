package problem0002

import "leetcode_go/types"

type ListNode = types.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    head := new(ListNode)
    list := head
    
    for l1!=nil || l2!=nil || carry!= 0{
        n1,n2:= 0,0
        if l1 != nil{
            n1,l1 = l1.Val, l1.Next
        }
        if l2 != nil{
            n2,l2 = l2.Val, l2.Next
        }
        currentsum := n1+n2+carry
        carry = (currentsum)/10
        list.Next = &ListNode{Val:currentsum%10}
        list = list.Next
    }
    
    return head.Next
    
}