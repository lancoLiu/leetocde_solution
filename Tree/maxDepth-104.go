package Tree

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/submissions/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归求解，
// 1、找返回值，
// 2、返回条件，
// 3、本级递归应该做什么逻辑操作
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
