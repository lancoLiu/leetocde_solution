package Tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

//给定一颗N叉树，返回节点值的前序遍历

var res4 []int

func preorder(root *Node) []int {
	res4 = []int{}
	dfsN(root)
	return res4

}

func dfsN(root *Node) {
	if root != nil {
		res4 = append(res4, root.Val)
		//遍历
		for _, value := range root.Children {
			dfsN(value)
		}

	}
}

//迭代解法

func preorder2(root *Node) []int {
	var r []int
	var stack []*Node
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}

	}
	return r
}
