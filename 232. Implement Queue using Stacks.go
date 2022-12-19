// https://leetcode.com/problems/implement-queue-using-stacks/description/

type MyQueue struct {
    inorder []int // push order
    reverse []int // reverse order
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (q *MyQueue) Push(x int)  {
    for len(q.reverse) != 0 { q.inorder = append(q.inorder, q.reverse[len(q.reverse)-1]); q.reverse = q.reverse[:len(q.reverse)-1] }
    q.inorder = append(q.inorder, x)
}

func (q *MyQueue) Pop() int {
    for len(q.inorder) != 0 { q.reverse = append(q.reverse, q.inorder[len(q.inorder)-1]); q.inorder = q.inorder[:len(q.inorder)-1] }
    pop := q.reverse[len(q.reverse)-1]
    q.reverse = q.reverse[:len(q.reverse)-1]
    return pop
}

func (q *MyQueue) Peek() int {
    if len(q.reverse) == 0 { return q.inorder[0] }
    return q.reverse[len(q.reverse)-1]
}

func (q *MyQueue) Empty() bool {
    return len(q.inorder) == 0 && len(q.reverse) == 0
}
