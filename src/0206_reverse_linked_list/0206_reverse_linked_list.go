package problem0206

import (
	"leetcode_go/types"
)

type ListNode = types.ListNode

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    for cur != nil{
        temp := cur.Next
        cur.Next = prev
        prev = cur
        cur = temp
    }
    return prev
}