package types

// ListNode is used for linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
// CreateListNode creates a list node
func CreateListNode(nums []int) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	for _, v := range nums {
		currentNode := &ListNode{Val: v, Next: nil}
		head.Next = currentNode
		head = head.Next
	}
	return head.Next
}
