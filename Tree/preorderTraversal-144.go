package Tree

var res1 []int

//递归解法
func preorderTraversal(root *TreeNode) []int {
	res1 = []int{}
	if root == nil {
		return nil
	}
	dfsPreorder(root)
	return res1
}

func dfsPreorder(root *TreeNode) {
	if root != nil {
		res1 = append(res1, root.Val)
		dfsPreorder(root.Left)
		dfsPreorder(root.Right)
	}

}

//迭代解法
//中左右
func PreorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val) //前序输出

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return res
}
