package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	//
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {
		tmpQ := []*TreeNode{}
		vals := []int{}

		for j := 0; j < len(queue); j++ {
			node := queue[j]
			vals = append(vals, node.Val)
			if node.Left != nil {
				tmpQ = append(tmpQ, node.Left)
			}
			if node.Right != nil {
				tmpQ = append(tmpQ, node.Right)
			}

		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if i%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		res = append(res, vals)
		queue = tmpQ
	}

	return res

}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	zigzagLevelOrder(tree)
}
