package Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//层序遍历再反转即可
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		tempQ := []*TreeNode{}
		tempR := []int{}
		//把当层的节点出栈存储到tempR中，并且将子节点值存储到tempQ中
		for i := 0; i < len(queue); i++ {
			cur := queue[i]
			tempR = append(tempR, cur.Val)
			if cur.Left != nil {
				tempQ = append(tempQ, cur.Left)
			}
			if cur.Right != nil {
				tempQ = append(tempQ, cur.Right)
			}
		}
		res = append(res, tempR)
		queue = tempQ
	}

	return reverse(res)
}
func reverse(s [][]int) [][]int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
