package main

import (
	"math"

	"lanco.vip/lib"
)

//平衡二叉树

//自顶向下递归
//root节点是否平衡且左右节点是否平衡
func topToBottom(root *lib.TreeNode) bool {

	if root == nil {
		return true
	}
	return math.Abs(float64(height(root.Left))-float64(height(root.Right))) <= 1 && topToBottom(root.Left) && topToBottom(root.Right)

}
func height(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func bottomToTop(root *lib.TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root) != -1
}

func recur(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	if math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}
	return max(left, right) + 1

}

func main() {
	t := &lib.TreeNode{
		Val: 4,
		Left: &lib.TreeNode{
			Val: 2,
			Left: &lib.TreeNode{
				Val: 1,
			},
			Right: &lib.TreeNode{
				Val: 3,
			},
		},
		Right: &lib.TreeNode{
			Val: 7,
			Left: &lib.TreeNode{
				Val: 6,
			},
			Right: &lib.TreeNode{
				Val: 9,
			},
		},
	}
	topToBottom(t)
}
