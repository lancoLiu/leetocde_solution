package main

import (
	"lanco.vip/lib"
)

//相同的树
//https://leetcode-cn.com/problems/same-tree/

func dfsSolve(p *lib.TreeNode, q *lib.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return dfsSolve(p.Left, q.Left) && dfsSolve(p.Right, q.Right)
}

//广度优先查找比较，非递归
func bfsSolve(p *lib.TreeNode, q *lib.TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	queue1 := []*lib.TreeNode{p}
	queue2 := []*lib.TreeNode{q}

	for len(queue1) > 0 && len(queue2) > 0 {
		node1, node2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if node1.Val != node2.Val {
			return false
		}
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if (left1 != nil && left2 == nil) || (left1 == nil && left2 != nil) {
			return false
		}
		if (right1 != nil && right2 == nil) || (right1 == nil && right2 != nil) {
			return false
		}

		if left1 != nil {
			queue1 = append(queue1, left1)
		}
		if right1 != nil {
			queue1 = append(queue1, right1)
		}
		if left2 != nil {
			queue2 = append(queue2, left2)
		}
		if right2 != nil {
			queue2 = append(queue2, right1)
		}
	}
	return len(queue1) == 0 && len(queue2) == 0
}

func main() {
	//test case
	//
}
