package problem0143

import (
	"leetcode_go/types"
)

type ListNode = types.ListNode

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil{
        return 
    }
    // find mid
    //ptr1 go to the mid, ptr2 go to the last
    ptr1, ptr2 := head, head
    for ptr2.Next != nil && ptr2.Next.Next != nil{
        ptr1 = ptr1.Next
        ptr2 = ptr2.Next.Next
    }
    
    // reverse list from mid to last
    ptrMid1 := ptr1.Next
    ptrMid2 := ptr1.Next.Next
    for ptrMid1.Next != nil{
        ptrMid1.Next = ptrMid2.Next
        ptrMid2.Next = ptr1.Next
        ptr1.Next = ptrMid2
        ptrMid2 = ptrMid1.Next
    }
    // fmt.Println(ptr1.Val, ptr1.Next.Val,ptr1.Next.Next.Val)
    // rearrange list from f
    ptrStart := head
    ptrMid := ptr1.Next
    for ptrStart != ptr1{
        ptr1.Next = ptrMid.Next
        ptrMid.Next = ptrStart.Next
        ptrStart.Next = ptrMid
        ptrStart = ptrMid.Next
        ptrMid = ptr1.Next
    }
}