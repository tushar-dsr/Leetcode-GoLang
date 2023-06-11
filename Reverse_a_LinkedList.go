// link: https://leetcode.com/problems/reverse-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//using extra space (stack)
func reverseList(head *ListNode) *ListNode {
	var buffer [](*ListNode) // used as a stack
	for head != nil {
		buffer = append(buffer, head)
		head = head.Next
	}
	if len(buffer) == 0 {
		return nil
	}
	head = buffer[len(buffer)-1]
	temp := head
	for i := len(buffer) - 2; i >= 0; i-- {
		temp.Next = buffer[i]
		temp = temp.Next
	}
	temp.Next = nil
	return head
}

// Using recurssion stack

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := reverse(head.Next)
	p := head.Next
	p.Next = head
	head.Next = nil
	return temp
}
func reverseList(head *ListNode) *ListNode {
	return reverse(head)
}
