package problem0142

import (
	"leetcode_go/types"
)

// ListNode is defined in leetcode_go/types
type ListNode = types.ListNode

func detectCycle(head *ListNode) *ListNode {
	//Floyd's algorithms is separated into two phases
	//1. determines whether a cycle exist, if not return nil
	//2. use the located intersection node to find entrance to the cycle
	if head == nil {
		return nil
	}
	//phase 1
	//slow = slow.Next, fast = fast.Next.Next, find intersection
	intersect := getIntersection(head)
	if intersect == nil {
		return nil
	}
	//phase 2
	//ptr1 = head, ptr2 = slow, advance each by one until they meet, return result
	ptr1 := head
	ptr2 := intersect
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}

func getIntersection(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
