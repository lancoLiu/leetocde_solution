package mai

import "lanco.vip/lib"

var res []int

//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
func kthSmallest(root *lib.TreeNode, k int) int {
	//先中序遍历存储在切片中，最后获取第k个元素即可
	res = []int{}
	dfs(root)
	return res[k-1]
}

func main() {

}

//中序遍历
func dfs(root *lib.TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	res = append(res, root.Val)
	dfs(root.Right)
}
