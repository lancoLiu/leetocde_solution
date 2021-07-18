package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	que := []*TreeNode{root}
	//最外层控制第一层数组元素
	for i := 0; len(que) > 0; i++ {
		ret = append(ret, []int{})
		//暂时切片，用来存储左右节点，后续替换为queue
		tempQ := []*TreeNode{}
		//第二层循环控制第二层数组元素
		for j := 0; j < len(que); j++ {
			//取出元素，但是没有去除元素，所以下标为j
			node := que[j]
			ret[i] = append(ret[i], node.Val)

			//存储该元素的左右元素
			if node.Left != nil {
				tempQ = append(tempQ, node.Left)
			}
			if node.Right != nil {
				tempQ = append(tempQ, node.Right)
			}

		}
		que = tempQ
	}
	return ret
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	levelOrder(tree)
}
