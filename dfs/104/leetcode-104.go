package main

import "lanco.vip/lib"

//二叉树最大深度

func maxDepth(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func bfsSolve(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*lib.TreeNode{root}
	ans := 0
	for len(queue) != 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}

		ans++
	}
	return ans

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
	maxDepth(t)
}
