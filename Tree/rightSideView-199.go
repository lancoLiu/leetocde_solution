package Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
站在右边看，所以我的想法是先层次遍历，再遍历结果集，如果只有一个值，就取第一个值；否则就是取最后一个值
*/
func rightSideView(root *TreeNode) []int {
	res := [][]int{}

	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tempQ := []*TreeNode{}
		tempR := []int{}

		for i := 0; i < len(queue); i++ {
			node := queue[i]
			tempR = append(tempR, node.Val)
			if node.Left != nil {
				tempQ = append(tempQ, node.Left)
			}
			if node.Right != nil {
				tempQ = append(tempQ, node.Right)
			}

		}
		res = append(res, tempR)
		queue = tempQ
	}
	r := []int{}

	for _, value := range res {
		if len(value) == 1 {
			r = append(r, value[0])
		} else {
			r = append(r, value[len(value)-1])
		}
	}

	return r
}
