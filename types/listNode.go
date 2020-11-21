package types

// ListNode is used for linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
// CreateListNode creates a list node
func CreateListNode(nums []int) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	p := head
	for _, v := range nums {
		currentNode := &ListNode{Val: v, Next: nil}
		p.Next = currentNode
		p = p.Next
	}
	return head.Next
}
