// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var total int
var product int
func maxProduct(root *TreeNode) int {
    total, product = 0, 0
    total = nodeSum(root)
    total = nodeSum(root)
    
    return product % (int(1e9) + 7)
}

func nodeSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    sum := root.Val + nodeSum(root.Left) + nodeSum(root.Right)
    
    product = max(product, (total - sum) * sum)
    
    return sum
}


func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}

