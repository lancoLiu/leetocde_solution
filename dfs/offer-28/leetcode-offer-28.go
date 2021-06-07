package main

import "lanco.vip/lib"

//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
func isSymmetric(root *lib.TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(a, b *lib.TreeNode) bool {
	//递归终止条件
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Right) && recur(a.Right, b.Left)
}

func main() {

}
