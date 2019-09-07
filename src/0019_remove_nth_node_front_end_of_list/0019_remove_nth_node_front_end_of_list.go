package problem0019

import "leetcode_go/types"

// ListNode is created from package types
type ListNode = types.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    //slow advance 1 node first time
    //fast advance n + 1 node first time
	//until the fast node go past the end of the list, remove slow.Next
	helper := &ListNode{Val : 0, Next : head}
	slow := helper
	fast := helper

	for i := 0; i <= n; i++{
		fast = fast.Next
	}

	for fast != nil{
		slow, fast = slow.Next, fast.Next
	}

	//remove target node
	slow.Next = slow.Next.Next

	return helper.Next 
}