package Tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
	2
null	3
	  null 4
		null 5
			null 6
再上图中最小深度不是1，我一开始以为是2->null的路径，但是2还有右边的子节点，所以他不是叶子结点
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/solution/111-er-cha-shu-de-zui-xiao-shen-du-di-gui-die-dai-/
为什么会犯这个错误，加入left==nil,right!=nil
说明最小深度是min(depth(right))+1

*/

func MinDepth(root *TreeNode) int {
	//结束条件
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(MinDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(MinDepth(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//错误解法

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
