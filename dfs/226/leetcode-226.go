package main

import (
	"lanco.vip/lib"
)

//翻转二叉树

func invertTree(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root

}

//迭代翻转
func bfsSolve(root *lib.TreeNode) *lib.TreeNode {
	if root == nil {
		return nil
	}

	queue := []*lib.TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
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
	bfsSolve(t)
	//test case
	//[4,2,7,1,3,6,9]
}
