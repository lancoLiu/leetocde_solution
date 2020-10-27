package Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归
var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	dfs(root)
	return res
}
func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
}

/**

     1
   2    7
  3
    4
  5    6
输出为[3 5 4 6 2 1]
*/
//非递归
func InorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		//直到将root的所有左节点都push入栈
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}
		curr := stack[len(stack)-1]
		res = append(res, curr.Val)  //中序输出
		root = curr.Right            //右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:len(stack)-1] //出栈
	}
	return res
}
