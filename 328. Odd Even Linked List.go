// https://leetcode.com/problems/odd-even-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var oddHead, oddTail, evenHead, evenTail *ListNode
	for current, i := head, 1; current != nil; current, i = current.Next, i+1 {
		if i%2 != 0 {
			if oddHead == nil {
				oddHead = current
				oddTail = oddHead
			} else {
				oddTail.Next = current
				oddTail = oddTail.Next
			}
		} else {
			if evenHead == nil {
				evenHead = current
				evenTail = evenHead
			} else {
				evenTail.Next = current
				evenTail = evenTail.Next
			}
		}
	}

	evenTail.Next = nil
	oddTail.Next = evenHead

	return oddHead
}
