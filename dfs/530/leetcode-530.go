package main

//给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

import (
	"math"

	"lanco.vip/lib"
)

func getMinimumDifference(root *lib.TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -1 && root.Val-pre < ans {
			ans = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return ans

}
func main() {
	//fmt.Println("hello")
}
