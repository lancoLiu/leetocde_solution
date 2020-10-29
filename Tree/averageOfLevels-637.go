package Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//先层序遍历，再遍历每层节点存储平均值

func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	res := [][]int{}

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
	r := []float64{}

	for _, value := range res {
		sum := 0.0
		nums := 0.0
		for _, v2 := range value {
			sum += float64(v2)
			nums++
		}
		r = append(r, sum/nums)
	}
	return r
}
