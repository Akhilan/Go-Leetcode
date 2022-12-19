// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    curr := head
    for curr.Next != nil {
        if (curr.Val == curr.Next.Val) {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return head;


}
