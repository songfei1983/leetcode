package middleofthelinkedlist

/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	var c int
	n := head.Next
	for {
		n = n.Next
		if n == nil {
			break
		}
		c++
	}
	n = head.Next
	for i := 0; i < c/2; i++ {
		n = n.Next
	}
	return n
}

// @lc code=end
