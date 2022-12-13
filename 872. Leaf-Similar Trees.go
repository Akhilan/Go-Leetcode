// https://leetcode.com/problems/leaf-similar-trees/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(node *TreeNode, leaf *([]int)) {
    if node == nil {
        return
    }
    if node.Left == nil && node.Right == nil {
        *leaf = append(*leaf, node.Val)
        return
    }
    dfs(node.Left, leaf)
    dfs(node.Right, leaf)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var t1, t2 []int
    dfs(root1, &t1)
    dfs(root2, &t2)
    if len(t1) != len(t2) {
        return false
    }
    for i, val := range t1 {
        if t2[i] != val {
            return false
        }
    }
    return true
}
