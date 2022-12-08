// https://leetcode.com/problems/middle-of-the-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    end, half := head, head
    
    for end != nil && end.Next != nil {
        
        half = half.Next
        end = end.Next.Next
        
    }
    return half
}
