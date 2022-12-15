// https://leetcode.com/problems/binary-tree-maximum-path-sum/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	out := -1000
	var dfs func(*TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		v, left, right := tn.Val, dfs(tn.Left), dfs(tn.Right)
		out = max(out, v, v+left, v+right, v+left+right)
		return max(v, v+left, v+right)
	}
	dfs(root)
	return out
}

func max(nums ...int) int {
	x := nums[0]
	for _, n := range nums {
		if n > x {
			x = n
		}
	}
	return x
}

