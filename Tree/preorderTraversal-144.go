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

//迭代解法P
func PreorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			res = append(res, root.Val)       //前序输出
			stack = append(stack, root.Right) //右节点 入栈
			root = root.Left                  //移至最左
		}
		index := len(stack) - 1 //栈顶
		root = stack[index]     //出栈
		stack = stack[:index]
	}
	return res
}
