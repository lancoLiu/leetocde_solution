package Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res3 []int

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res3 = []int{}
	dfsPostorder(root)
	return res3

}

func dfsPostorder(root *TreeNode) {
	if root != nil {
		dfsPostorder(root.Left)
		dfsPostorder(root.Right)
		res3 = append(res3, root.Val)
	}
}

//迭代解法
/**
左右中
中左右
  右左中（通过切片复制操作）
  左右中（先压左子树）
*/
func PostorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	//stack, rest := Stack([]*TreeNode{root}), []int{}
	res := []int{}
	var stack []*TreeNode
	stack = append(stack, root)
	//421
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//这一步是关键
		res = append([]int{curr.Val}, res...)
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}
	return res

}
