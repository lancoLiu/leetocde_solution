package main

import "lanco.vip/lib"

//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

func mergeTrees(root1 *lib.TreeNode, root2 *lib.TreeNode) *lib.TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func main() {

}
