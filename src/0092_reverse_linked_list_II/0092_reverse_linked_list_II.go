package problem0092

import (
	"leetcode_go/types"
)

type ListNode = types.ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    shadowHead := &ListNode{Next:head}
    start := shadowHead
    for i :=1; i < m; i++{
        start = start.Next
    }
	p1 := start.Next
	p2 := start.Next.Next
    for i := m; i < n; i++{
        p1.Next = p2.Next
        p2.Next = start.Next
        start.Next = p2
        p2 = p1.Next
    }
    return shadowHead.Next
}