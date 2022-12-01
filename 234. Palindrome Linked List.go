// https://leetcode.com/problems/palindrome-linked-list/


// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.


func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	cur := slow
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}
	return true
}
