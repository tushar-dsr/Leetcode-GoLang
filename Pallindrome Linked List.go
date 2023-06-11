// link: https://leetcode.com/problems/palindrome-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//1st solution
func isPalindrome(head *ListNode) bool {
	buffer := []int{}
	for head != nil {
		buffer = append(buffer, head.Val)
		head = head.Next
	}
	for i, j := 0, len(buffer)-1; i < j; i, j = i+1, j-1 {
		if buffer[i] != buffer[j] {
			return false
		}
	}
	return true
}

