package main

import "lanco.vip/lib"

//如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
func isUnivalTree(root *lib.TreeNode) bool {

	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)

}

func main() {

}
