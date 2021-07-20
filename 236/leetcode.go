package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
/**
1、终止条件：a.遇到节点为p或者q的;b.到达叶子节点了
2、如果left为空，right不为空
3、
4、
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//这里比较难想
	//当left为空，说明p,q都不在左子树中
	//讨论：
	//1、p,q 其中一个在root的右子树中，此时 right指向p(假设为p)；
	//2、p,q 两节点都在 root的右子树中，此时的right指向最近公共祖先节点 ；
	//3、两个都不在root的左右子树中，向上返回
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	//如果上述都不为空，说明当前root就是最近的祖先节点，可以返回
	return root
}
