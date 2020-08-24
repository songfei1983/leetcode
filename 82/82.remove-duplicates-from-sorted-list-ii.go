package removeduplicatesfromsortedlistii

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	first := head
	second := head
	for first != nil {
		if first.Next != nil && first.Val == first.Next.Val {
			second = first.Next
			for second.Next != nil && second.Val == second.Next.Val {
				second = second.Next
			}
			first = second.Next
			if first == nil {
				prev.Next = nil
			}
			continue
		}
		prev.Next = first
		first = first.Next
		prev = prev.Next
	}
	return dummy.Next
}

// @lc code=end
