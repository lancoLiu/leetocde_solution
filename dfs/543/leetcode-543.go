package main

import "lanco.vip/lib"

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
var maxV int

func diameterOfBinaryTree(root *lib.TreeNode) int {
	maxV = 0

	dfs(root)
	return maxV - 1
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	maxV = max(maxV, left+right+1)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

}
