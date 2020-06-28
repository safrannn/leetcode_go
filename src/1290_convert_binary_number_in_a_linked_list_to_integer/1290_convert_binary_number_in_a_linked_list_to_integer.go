package problem1290

import "leetcode_go/types"

type ListNode = types.ListNode

func getDecimalValue(head *ListNode) int {  
    result := 0
    for head != nil{
        result = result << 1 + head.Val;
        head = head.Next
    }
    return result
}