// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

    if root.Val < p.Val && root.Val < q.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else if root.Val > p.Val && root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }


    return root
}



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

//Recursive

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

  for {
         if root.Val>p.Val && root.Val>q.Val{
             root=root.Left
         } else if root.Val<p.Val && root.Val<q.Val{
             root=root.Right
         } else {
             return root
         }
     }
}
