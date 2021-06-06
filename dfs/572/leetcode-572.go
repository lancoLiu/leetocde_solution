package main

import "lanco.vip/lib"

//给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

func isSubtree(root *lib.TreeNode, subRoot *lib.TreeNode) bool {
	//必须root和subroot同时都到达nil，才返回true
	return (root != nil && subRoot != nil) && (recur(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot))
}

func recur(t1 *lib.TreeNode, t2 *lib.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}
	return recur(t1.Left, t2.Left) && recur(t1.Right, t2.Right)
}

func main() {

}
