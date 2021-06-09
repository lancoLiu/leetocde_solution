package main

import "lanco.vip/lib"

func convertBST(root *lib.TreeNode) *lib.TreeNode {
	//中序遍历的倒序即可
	var dfs func(root *lib.TreeNode)
	ans := 0
	dfs = func(node *lib.TreeNode) {
		if node != nil {
			dfs(node.Right)
			ans += node.Val
			node.Val = ans
			dfs(node.Left)
		}

	}
	dfs(root)
	return root

}

func main() {

}
